package model

import (
	"github.com/google/uuid"
	"loanManagement/constant"
	databaseModel "loanManagement/database/model"
)

type FetchAdminLoansInput struct {
	LoanIDs                 []uuid.UUID
	Status                  *constant.LoanStatus
	FetchScheduledRepayment bool
}

type FetchAdminLoansOutput struct {
	Loans                   []*databaseModel.Loan
	LoanScheduledRepayments map[uuid.UUID][]*databaseModel.ScheduledRepayment
}

type UpdateLoanAndScheduledRepaymentInput struct {
	LoanI              *databaseModel.Loan
	ScheduleRepayments []*databaseModel.ScheduledRepayment
	Status             constant.LoanStatus
}

type UpdateLoanAndScheduledRepaymentOutput struct {
	Success bool
}
