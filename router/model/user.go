package model

import (
	"time"

	"loanManagement/constant"

	"github.com/google/uuid"
)

type UserLoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginOutput struct {
	Token string `json:"token"`
}

type UserCreateLoanInput struct {
	Amount        int64             `json:"amount"`
	Currency      constant.Currency `json:"currency"`
	Term          int64             `json:"term"` // in weeks
	DisbursalDate time.Time         `json:"disbursalDate"`
}

type UserCreateLoanOutput struct {
	LoanID uuid.UUID `json:"loanID"`
}
