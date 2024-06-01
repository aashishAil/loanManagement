package model

import (
	"gorm.io/gorm"
	"time"

	"loanManagement/constant"

	"github.com/google/uuid"
)

type BulkCreateScheduledRepaymentInput struct {
	LoanID          uuid.UUID
	ScheduledAmount int64
	Currency        constant.Currency
	ScheduledDates  []time.Time
	TxDb            *gorm.DB
}

type UpdateScheduledRepaymentInput struct {
	ID            uuid.UUID
	PendingAmount *int64
	Status        *constant.SchedulePaymentStatus
	TxDb          *gorm.DB
}

type FindAllScheduledRepaymentInput struct {
	LoanID uuid.UUID
}
