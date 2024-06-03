package model

import (
	"time"

	"loanManagement/constant"

	"github.com/google/uuid"
)

type CreateUserLoanInput struct {
	UserID        uuid.UUID
	Amount        int64
	Currency      constant.Currency
	Term          int64 // in weeks
	DisbursalDate time.Time
}
