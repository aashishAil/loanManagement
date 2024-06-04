package handler

import (
	"context"

	"loanManagement/appError"
	"loanManagement/constant"
	dbInstance "loanManagement/database/instance"
	databaseModel "loanManagement/database/model"
	handlerModel "loanManagement/handler/model"
	"loanManagement/logger"
	"loanManagement/repo"
	repoModel "loanManagement/repo/model"

	"github.com/google/uuid"
)

type Admin interface {
	FetchLoans(ctx context.Context, data handlerModel.FetchAdminLoansInput) (*handlerModel.FetchAdminLoansOutput, error)
	UpdateLoanAndScheduledRepayment(ctx context.Context, data handlerModel.UpdateLoanAndScheduledRepaymentInput) (*handlerModel.UpdateLoanAndScheduledRepaymentOutput, error)
}

type admin struct {
	loanRepo               repo.Loan
	scheduledRepaymentRepo repo.ScheduledRepayment

	dbInstance dbInstance.PostgresDB
}

func (h *admin) FetchLoans(ctx context.Context, data handlerModel.FetchAdminLoansInput) (*handlerModel.FetchAdminLoansOutput, error) {
	loansArr, err := h.loanRepo.FindAll(ctx, repoModel.FindAllLoanInput{
		IDs:    data.LoanIDs,
		Status: data.Status,
	})
	if err != nil {
		logger.Log.Error("failed to find loans", logger.Error(err))
		return nil, err
	}

	loanScheduledRepayments := make(map[uuid.UUID][]*databaseModel.ScheduledRepayment)

	if len(loansArr) == 0 || data.FetchScheduledRepayment == false {
		// no loans found for the user
		return &handlerModel.FetchAdminLoansOutput{
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

	return &handlerModel.FetchAdminLoansOutput{
		Loans:                   loansArr,
		LoanScheduledRepayments: loanScheduledRepayments,
	}, nil
}

func (h *admin) UpdateLoanAndScheduledRepayment(ctx context.Context, data handlerModel.UpdateLoanAndScheduledRepaymentInput) (*handlerModel.UpdateLoanAndScheduledRepaymentOutput, error) {
	if data.LoanI.Status == constant.LoanStatusPaid || data.LoanI.Status == constant.LoanStatusRejected ||
		data.LoanI.Status == constant.LoanStatusApproved {
		return nil, appError.Custom{
			Err: constant.LoanInTerminalStatusError,
		}
	}
	scheduleRepaymentStatus := constant.ScheduleRepaymentStatusApproved
	if data.Status == constant.LoanStatusRejected {
		scheduleRepaymentStatus = constant.ScheduleRepaymentStatusRejected
	}

	txnDb, err := h.dbInstance.GetTransactionDb()
	if err != nil {
		logger.Log.Error("failed to start transaction", logger.Error(err))
		return nil, err
	}

	err = h.loanRepo.Update(ctx, repoModel.UpdateLoanInput{
		ID:     data.LoanI.ID,
		Status: &data.Status,
		TxDb:   &txnDb,
	})
	if err != nil {
		logger.Log.Error("failed to update loan", logger.Error(err))
		rollbackError := txnDb.Rollback()
		if rollbackError != nil {
			logger.Log.Error("failed to rollback transaction", logger.Error(rollbackError))
			return nil, rollbackError
		}
		return nil, err
	}

	err = h.scheduledRepaymentRepo.Update(ctx, repoModel.UpdateScheduledRepaymentInput{
		LoanID: &data.LoanI.ID,
		Status: &scheduleRepaymentStatus,
		TxDb:   &txnDb,
	})
	if err != nil {
		logger.Log.Error("failed to update scheduled repayments", logger.Error(err))
		rollbackError := txnDb.Rollback()
		if rollbackError != nil {
			logger.Log.Error("failed to rollback transaction", logger.Error(rollbackError))
			return nil, rollbackError
		}
		return nil, err
	}

	err = txnDb.Commit()
	if err != nil {
		logger.Log.Error("failed to commit transaction", logger.Error(err))
		return nil, err
	}

	logger.Log.Info("updated loan and scheduled repayments",
		logger.String("loanID", data.LoanI.ID.String()))
	return &handlerModel.UpdateLoanAndScheduledRepaymentOutput{
		Success: true,
	}, nil
}

func NewAdmin(
	loanRepo repo.Loan,
	scheduledRepaymentRepo repo.ScheduledRepayment,

	dbInstance dbInstance.PostgresDB,
) Admin {
	return &admin{
		loanRepo:               loanRepo,
		scheduledRepaymentRepo: scheduledRepaymentRepo,

		dbInstance: dbInstance,
	}
}
