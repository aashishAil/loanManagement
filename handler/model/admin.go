package model

import (
	"github.com/google/uuid"
	"loanManagement/constant"
	databaseModel "loanManagement/database/model"
)

type FetchAdminLoansInput struct {
	Status                  constant.LoanStatus
	FetchScheduledRepayment bool
}

type FetchAdminLoansOutput struct {
	Loans                   []*databaseModel.Loan
	LoanScheduledRepayments map[uuid.UUID][]*databaseModel.ScheduledRepayment
}
