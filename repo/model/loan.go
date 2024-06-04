package model

import (
	"time"

	"loanManagement/constant"
	dbInstance "loanManagement/database/instance"

	"github.com/google/uuid"
)

type CreateLoanInput struct {
	UserID        uuid.UUID
	Amount        int64
	Currency      constant.Currency
	Term          int64
	DisbursalDate time.Time
	TxDb          *dbInstance.PostgresTransactionDB
}

type FindOneLoanInput struct {
	ID            *uuid.UUID
	UserID        *uuid.UUID
	DisbursalDate *time.Time
}

type FindAllLoanInput struct {
	IDs    []uuid.UUID
	UserID *uuid.UUID
	Status *constant.LoanStatus
}

type UpdateLoanInput struct {
	ID            uuid.UUID
	Status        *constant.LoanStatus
	PendingAmount *int64
	TxDb          *dbInstance.PostgresTransactionDB
}
