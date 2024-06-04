package handler

import (
	"context"

	databaseModel "loanManagement/database/model"
	handlerModel "loanManagement/handler/model"
	"loanManagement/logger"
	"loanManagement/repo"
	repoModel "loanManagement/repo/model"

	"github.com/google/uuid"
)

type Admin interface {
	FetchLoans(ctx context.Context, data handlerModel.FetchAdminLoansInput) (*handlerModel.FetchAdminLoansOutput, error)
}

type admin struct {
	loanRepo               repo.Loan
	scheduledRepaymentRepo repo.ScheduledRepayment
}

func (h *admin) FetchLoans(ctx context.Context, data handlerModel.FetchAdminLoansInput) (*handlerModel.FetchAdminLoansOutput, error) {
	loansArr, err := h.loanRepo.FindAll(ctx, repoModel.FindAllLoanInput{
		Status: &data.Status,
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

func NewAdmin(
	loanRepo repo.Loan,
	scheduledRepaymentRepo repo.ScheduledRepayment,
) Admin {
	return &admin{
		loanRepo:               loanRepo,
		scheduledRepaymentRepo: scheduledRepaymentRepo,
	}
}
