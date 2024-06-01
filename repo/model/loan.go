package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"loanManagement/constant"
	"time"
)

type CreateLoanInput struct {
	UserID        uuid.UUID
	Amount        int64
	Currency      constant.Currency
	Term          int64
	DisbursalDate time.Time
	TxDb          *gorm.DB
}

type FindOneLoanInput struct {
	UserID        uuid.UUID
	DisbursalDate *time.Time
}

type FindAllLoanInput struct {
	UserID uuid.UUID
}

type UpdateLoanInput struct {
	ID              uuid.UUID
	Status          constant.LoanStatus
	DisbursalAmount *int64
	DisbursalDate   *time.Time
	TxDb            *gorm.DB
}
