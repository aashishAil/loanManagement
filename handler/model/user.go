package model

import (
	"time"

	"loanManagement/constant"
	databaseModel "loanManagement/database/model"

	"github.com/google/uuid"
)

type CreateUserLoanInput struct {
	UserID        uuid.UUID
	Amount        int64
	Currency      constant.Currency
	Term          int64 // in weeks
	DisbursalDate time.Time
}

type FetchUserLoanInput struct {
	UserID uuid.UUID
}

type FetchUserLoansOutput struct {
	Loans                   []*databaseModel.Loan
	LoanScheduledRepayments map[uuid.UUID][]*databaseModel.ScheduledRepayment
}
