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
	AddLoanPayment(ctx context.Context, data handlerModel.AddUserLoanPaymentInput) (*handlerModel.AddUserLoanPaymentOutput, error)
}

type user struct {
	loanRepo               repo.Loan
	paymentRepo            repo.Payment
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

func (h *user) AddLoanPayment(ctx context.Context, data handlerModel.AddUserLoanPaymentInput) (*handlerModel.AddUserLoanPaymentOutput, error) {
	loanI, err := h.loanRepo.FindOne(ctx, repoModel.FindOneLoanInput{
		ID:     &data.LoanID,
		UserID: &data.UserID,
	})
	if err != nil {
		logger.Log.Error("failed to find loan", logger.Error(err))
		return nil, err
	}

	if loanI.Status == constant.LoanStatusPending || loanI.Status == constant.LoanStatusPaid || loanI.Status == constant.LoanStatusRejected {
		logger.Log.Info("invalid loan status for payment", logger.String("status", string(loanI.Status)))
		return nil, appError.Custom{
			Err:  constant.LoanInTerminalStatusError,
			Code: http.StatusBadRequest,
		}
	}

	var nextScheduledRepayment *databaseModel.ScheduledRepayment
	markLoanAsPaid := false
	paidAmount := int64(data.Amount * constant.MinCurrencyConversionFactor)
	pendingAmount := loanI.PendingAmount - paidAmount
	if pendingAmount < 0 {
		logger.Log.Info("invalid payment amount", logger.String("loanID", loanI.ID.String()),
			logger.Float64("amount", data.Amount))
		return nil, appError.Custom{
			Err:  constant.InvalidPaymentAmountError,
			Code: http.StatusBadRequest,
		}
	}

	if pendingAmount == 0 {
		markLoanAsPaid = true
	}

	scheduledRepaymentStatus := constant.ScheduleRepaymentStatusApproved
	scheduledRepayments, err := h.scheduledRepaymentRepo.FindAll(ctx, repoModel.FindAllScheduledRepaymentInput{
		LoanIDs: []uuid.UUID{data.LoanID},
		Status:  &scheduledRepaymentStatus,
	})
	if err != nil {
		logger.Log.Error("failed to find scheduled repayments", logger.String("loanID", loanI.ID.String()), logger.Error(err))
		return nil, err
	}

	if len(scheduledRepayments) == 0 {
		logger.Log.Error("no pending scheduled repayments found for loan", logger.String("loanID", loanI.ID.String()), logger.String("loanID", data.LoanID.String()))
		return nil, appError.Custom{
			Err:  constant.NoScheduledRepaymentFoundError,
			Code: http.StatusInternalServerError,
		}
	}

	if paidAmount < scheduledRepayments[0].PendingAmount {
		return nil, appError.Custom{
			Err:  constant.PaymentAmountScheduledPaymentMismatchError,
			Code: http.StatusBadRequest,
		}
	}

	txnDb, err := h.dbInstance.GetTransactionDb()
	if err != nil {
		logger.Log.Error("failed to create transaction db", logger.String("loanID", loanI.ID.String()), logger.Error(err))
		return nil, err
	}

	_, err = h.paymentRepo.Create(ctx, repoModel.CreatePaymentInput{
		LoanID:   data.LoanID,
		UserID:   data.UserID,
		Amount:   data.Amount,
		Currency: loanI.Currency,
		TxDb:     &txnDb,
	})
	if err != nil {
		rollbackErr := txnDb.Rollback()
		if rollbackErr != nil {
			logger.Log.Error("failed to rollback transaction", logger.String("loanID", loanI.ID.String()), logger.Error(rollbackErr))
			return nil, rollbackErr
		}
		return nil, err
	}

	updateLoanInput := repoModel.UpdateLoanInput{
		ID:            data.LoanID,
		PendingAmount: &pendingAmount,
	}
	if markLoanAsPaid {
		loanStatus := constant.LoanStatusPaid
		updateLoanInput.Status = &loanStatus
	}

	err = h.loanRepo.Update(ctx, updateLoanInput)
	if err != nil {
		rollbackErr := txnDb.Rollback()
		if rollbackErr != nil {
			logger.Log.Error("failed to rollback transaction", logger.String("loanID", loanI.ID.String()), logger.Error(rollbackErr))
			return nil, rollbackErr
		}
		return nil, err
	}

	for i := range scheduledRepayments {
		scheduledRepaymentI := scheduledRepayments[i]
		markScheduledRepaymentAsPaid := false
		if paidAmount <= 0 {
			nextScheduledRepayment = scheduledRepaymentI
			break
		}
		scheduledPaymentAmount := scheduledRepaymentI.PendingAmount - paidAmount
		if scheduledPaymentAmount < 0 {
			scheduledPaymentAmount = 0
		}
		if scheduledPaymentAmount == 0 {
			markScheduledRepaymentAsPaid = true
		}
		paidAmount -= scheduledRepaymentI.PendingAmount
		updateScheduledRepaymentInput := repoModel.UpdateScheduledRepaymentInput{
			ID:            &scheduledRepaymentI.ID,
			PendingAmount: &scheduledPaymentAmount,
		}
		if markScheduledRepaymentAsPaid {
			scheduledRepaymentStatus := constant.ScheduleRepaymentStatusPaid
			updateScheduledRepaymentInput.Status = &scheduledRepaymentStatus
		}
		err = h.scheduledRepaymentRepo.Update(ctx, updateScheduledRepaymentInput)
		if err != nil {
			rollbackErr := txnDb.Rollback()
			if rollbackErr != nil {
				logger.Log.Error("failed to rollback transaction", logger.String("loanID", loanI.ID.String()), logger.Error(rollbackErr))
				return nil, rollbackErr
			}
			return nil, err
		}
		if scheduledPaymentAmount > 0 && paidAmount == 0 {
			scheduledRepaymentI.PendingAmount = scheduledPaymentAmount
			nextScheduledRepayment = scheduledRepaymentI
			break
		}
	}

	err = txnDb.Commit()
	if err != nil {
		logger.Log.Error("failed to commit transaction", logger.Error(err))
		return nil, err
	}

	resp := &handlerModel.AddUserLoanPaymentOutput{
		IsLoanClosed:  markLoanAsPaid,
		PendingAmount: float64(pendingAmount) / constant.MinCurrencyConversionFactor,
	}
	if nextScheduledRepayment != nil {
		pendingScheduledRepaymentAmount := float64(nextScheduledRepayment.PendingAmount) / constant.MinCurrencyConversionFactor
		resp.NextPaymentAmount = &pendingScheduledRepaymentAmount
		resp.NextDueDate = &nextScheduledRepayment.ScheduledDate
	}
	return resp, nil
}

func NewUser(
	loanRepo repo.Loan,
	paymentRepo repo.Payment,
	scheduledRepaymentRepo repo.ScheduledRepayment,
	userRepo repo.User,

	dbInstance dbInstance.PostgresDB,

	jwtUtil util.Jwt,
	timeUtil util.Time,
) User {
	return &user{
		loanRepo:               loanRepo,
		paymentRepo:            paymentRepo,
		scheduledRepaymentRepo: scheduledRepaymentRepo,
		userRepo:               userRepo,

		dbInstance: dbInstance,

		jwtUtil:  jwtUtil,
		timeUtil: timeUtil,
	}
}
