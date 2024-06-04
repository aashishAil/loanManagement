package model

import (
	dbInstance "loanManagement/database/instance"
	"time"

	"loanManagement/constant"

	"github.com/google/uuid"
)

type BulkCreateScheduledRepaymentInput struct {
	LoanID         uuid.UUID
	LoanAmount     int64
	Currency       constant.Currency
	ScheduledDates []time.Time
	TxDb           *dbInstance.PostgresTransactionDB
}

type UpdateScheduledRepaymentInput struct {
	ID            *uuid.UUID
	LoanID        *uuid.UUID
	PendingAmount *int64
	Status        *constant.ScheduleRepaymentStatus
	TxDb          *dbInstance.PostgresTransactionDB
}

type FindAllScheduledRepaymentInput struct {
	LoanIDs []uuid.UUID
	Status  *constant.ScheduleRepaymentStatus
}
