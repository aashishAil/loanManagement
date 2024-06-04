package handler

import (
	"context"
	"net/http"
	"time"

	"loanManagement/appError"
	"loanManagement/constant"
	dbInstance "loanManagement/database/instance"
	databaseModel "loanManagement/database/model"
	handlerModel "loanManagement/handler/model"
	"loanManagement/logger"
	"loanManagement/repo"
	repoModel "loanManagement/repo/model"
	"loanManagement/util"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type User interface {
	CheckValidCredentials(ctx context.Context, email, password string) (string, error)
	CreateLoan(ctx context.Context, data handlerModel.CreateUserLoanInput) (uuid.UUID, error)
	FetchLoans(ctx context.Context, data handlerModel.FetchUserLoanInput) (*handlerModel.FetchUserLoansOutput, error)
}

type user struct {
	loanRepo               repo.Loan
	scheduledRepaymentRepo repo.ScheduledRepayment
	userRepo               repo.User

	dbInstance dbInstance.PostgresDB

	jwtUtil  util.Jwt
	timeUtil util.Time
}

func (h *user) CheckValidCredentials(ctx context.Context, email, password string) (string, error) {
	userI, err := h.userRepo.FindOne(ctx, repoModel.FindOneUserInput{
		Email:    email,
		Password: password,
	})
	if err != nil {
		customErr := appError.Custom{}
		if errors.As(err, &customErr) {
			return "", appError.Custom{
				Err:  customErr.Err,
				Code: http.StatusUnauthorized,
			}
		}
		logger.Log.Error("failed to find user", logger.Error(err))
		return "", appError.Custom{
			Err:  errors.New("failed to find user"),
			Code: http.StatusInternalServerError,
		}
	}

	if userI == nil {
		logger.Log.Info("user not found", logger.String("email", email))
		return "", appError.Custom{
			Err:  errors.New("user not found"),
			Code: http.StatusNotFound,
		}
	}

	token, err := h.jwtUtil.GenerateToken(*userI)
	if err != nil {
		logger.Log.Error("failed to generate token", logger.Error(err))
		return "", appError.Custom{
			Err:  errors.New("failed to generate token"),
			Code: http.StatusInternalServerError,
		}
	}

	return token, nil
}

func (h *user) CreateLoan(ctx context.Context, data handlerModel.CreateUserLoanInput) (uuid.UUID, error) {

	// check if DisbursalDate is greater than 7 days from current time
	if data.DisbursalDate.Before(h.timeUtil.GetCurrent().AddDate(0, 0, constant.LoanCreationDisbursalTimeGap)) {
		logger.Log.Info("invalid disbursal date gap",
			logger.String("disbursalDate", data.DisbursalDate.String()),
			logger.String("currentDate", h.timeUtil.GetCurrent().String()),
		)
		return uuid.Nil, appError.Custom{
			Err:  constant.InvalidDisbursalDateGapError,
			Code: http.StatusBadRequest,
		}
	}

	if data.Term < constant.MinLoanRepaymentTenure {
		logger.Log.Info("term is less than minimum", logger.Int64("term", data.Term))
		return uuid.Nil, appError.Custom{
			Err:  constant.MinRepaymentTenureError,
			Code: http.StatusBadRequest,
		}

	}

	if data.Term > constant.MaxLoanRepaymentTenure {
		logger.Log.Info("term is greater than maximum", logger.Int64("term", data.Term))
		return uuid.Nil, appError.Custom{
			Err:  constant.MaxRepaymentTenureError,
			Code: http.StatusBadRequest,
		}
	}

	if data.Amount < constant.MinDisbursalAmount {
		logger.Log.Info("amount is less than minimum", logger.Int64("amount", data.Amount))
		return uuid.Nil, appError.Custom{
			Err:  constant.InvalidMinAmountError,
			Code: http.StatusBadRequest,
		}
	}

	if data.Amount > constant.MaxDisbursalAmount {
		logger.Log.Info("amount is greater than maximum", logger.Int64("amount", data.Amount))
		return uuid.Nil, appError.Custom{
			Err:  constant.InvalidMaxAmountError,
			Code: http.StatusBadRequest,
		}
	}

	if _, found := constant.CurrencyMap[data.Currency]; !found {
		logger.Log.Info("invalid currency", logger.String("currency", string(data.Currency)))
		return uuid.Nil, appError.Custom{
			Err:  constant.InvalidCurrencyError,
			Code: http.StatusBadRequest,
		}
	}

	txnDb, err := h.dbInstance.GetTransactionDb()
	if err != nil {
		return uuid.Nil, err
	}

	loanI, err := h.loanRepo.Create(ctx, repoModel.CreateLoanInput{
		UserID:        data.UserID,
		Amount:        data.Amount,
		Currency:      data.Currency,
		Term:          data.Term,
		DisbursalDate: data.DisbursalDate,
		TxDb:          &txnDb,
	})

	if err != nil {
		rollbackErr := txnDb.Rollback()
		if rollbackErr != nil {
			logger.Log.Error("failed to rollback transaction", logger.Error(rollbackErr))
			return uuid.Nil, errors.Wrap(err, "failed to rollback transaction")
		}
		return uuid.Nil, err
	}

	repaymentDates := make([]time.Time, data.Term)
	for i := 0; i < int(data.Term); i++ {
		repaymentDates[i] = data.DisbursalDate.AddDate(0, 0, (i+1)*constant.LoanRepaymentTimeGap)
	}

	err = h.scheduledRepaymentRepo.BulkCreate(ctx, repoModel.BulkCreateScheduledRepaymentInput{
		LoanID:         loanI.ID,
		LoanAmount:     data.Amount,
		Currency:       data.Currency,
		ScheduledDates: repaymentDates,
		TxDb:           &txnDb,
	})
	if err != nil {
		rollbackErr := txnDb.Rollback()
		if rollbackErr != nil {
			logger.Log.Error("failed to rollback transaction", logger.Error(rollbackErr))
			return uuid.Nil, errors.Wrap(err, "failed to rollback transaction")
		}
		return uuid.Nil, err
	}

	err = txnDb.Commit()
	if err != nil {
		logger.Log.Error("failed to commit transaction", logger.Error(err))
		return uuid.Nil, errors.Wrap(err, "failed to commit transaction")
	}

	return loanI.ID, nil
}

func (h *user) FetchLoans(ctx context.Context, data handlerModel.FetchUserLoanInput) (*handlerModel.FetchUserLoansOutput, error) {
	loansArr, err := h.loanRepo.FindAll(ctx, repoModel.FindAllLoanInput{
		UserID: &data.UserID,
	})
	if err != nil {
		logger.Log.Error("failed to find loans", logger.Error(err))
		return nil, err
	}

	loanScheduledRepayments := make(map[uuid.UUID][]*databaseModel.ScheduledRepayment)

	if len(loansArr) == 0 {
		// no loans found for the user
		return &handlerModel.FetchUserLoansOutput{
			Loans:                   loansArr,
			LoanScheduledRepayments: loanScheduledRepayments,
		}, nil
	}

	loanIDs := make([]uuid.UUID, len(loansArr))
	for i := range loansArr {
		loanI := loansArr[i]
		loanIDs[i] = loanI.ID
	}

	scheduledRepayments, err := h.scheduledRepaymentRepo.FindAll(ctx, repoModel.FindAllScheduledRepaymentInput{
		LoanIDs: loanIDs,
	})
	if err != nil {
		logger.Log.Error("failed to find scheduled repayments", logger.Error(err))
		return nil, err
	}

	for i := range scheduledRepayments {
		scheduledRepaymentI := scheduledRepayments[i]
		loanScheduledRepayments[scheduledRepaymentI.LoanID] = append(loanScheduledRepayments[scheduledRepaymentI.LoanID], scheduledRepaymentI)
	}

	return &handlerModel.FetchUserLoansOutput{
		Loans:                   loansArr,
		LoanScheduledRepayments: loanScheduledRepayments,
	}, nil
}

func NewUser(
	loanRepo repo.Loan,
	scheduledRepaymentRepo repo.ScheduledRepayment,
	userRepo repo.User,

	dbInstance dbInstance.PostgresDB,

	jwtUtil util.Jwt,
	timeUtil util.Time,
) User {
	return &user{
		loanRepo:               loanRepo,
		scheduledRepaymentRepo: scheduledRepaymentRepo,
		userRepo:               userRepo,

		dbInstance: dbInstance,

		jwtUtil:  jwtUtil,
		timeUtil: timeUtil,
	}
}
