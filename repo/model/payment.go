package model

import (
	"gorm.io/gorm"
	"loanManagement/constant"

	"github.com/google/uuid"
)

type CreatePaymentInput struct {
	LoanID   uuid.UUID
	UserID   uuid.UUID
	Amount   int64
	Currency constant.Currency
	TxDb     *gorm.DB
}
