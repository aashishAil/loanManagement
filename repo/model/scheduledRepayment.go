package model

import (
	dbInstance "loanManagement/database/instance"
	"time"

	"loanManagement/constant"

	"github.com/google/uuid"
)

type BulkCreateScheduledRepaymentInput struct {
	LoanID          uuid.UUID
	ScheduledAmount int64
	Currency        constant.Currency
	ScheduledDates  []time.Time
	TxDb            *dbInstance.PostgresTransactionDB
}

type UpdateScheduledRepaymentInput struct {
	ID            uuid.UUID
	PendingAmount *int64
	Status        *constant.SchedulePaymentStatus
	TxDb          *dbInstance.PostgresTransactionDB
}

type FindAllScheduledRepaymentInput struct {
	LoanID uuid.UUID
}
