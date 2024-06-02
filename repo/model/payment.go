package model

import (
	"loanManagement/constant"
	dbInstance "loanManagement/database/instance"

	"github.com/google/uuid"
)

type CreatePaymentInput struct {
	LoanID   uuid.UUID
	UserID   uuid.UUID
	Amount   int64
	Currency constant.Currency
	TxDb     *dbInstance.PostgresTransactionDB
}
