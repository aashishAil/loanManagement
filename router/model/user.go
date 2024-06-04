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

type GetUserLoansOutput struct {
	Loans []UserLoan `json:"loan"`
}

type UserLoan struct {
	ID                  uuid.UUID                 `json:"id"`
	UserID              uuid.UUID                 `json:"userID"`
	DisbursalAmount     float64                   `json:"disbursalAmount"`
	PendingAmount       float64                   `json:"pendingAmount"`
	Currency            constant.Currency         `json:"currency"`
	Term                int64                     `json:"term"`
	Status              constant.LoanStatus       `json:"status"`
	DisbursalDate       time.Time                 `json:"disbursalDate"`
	ScheduledRepayments *[]UserScheduledRepayment `json:"scheduledRepayments"`
}

type UserScheduledRepayment struct {
	ID              uuid.UUID                      `json:"id"`
	LoanID          uuid.UUID                      `json:"loanID"`
	ScheduledAmount float64                        `json:"scheduledAmount"`
	PendingAmount   float64                        `json:"pendingAmount"`
	Currency        constant.Currency              `json:"currency"`
	Status          constant.SchedulePaymentStatus `json:"status"`
	ScheduledDate   time.Time                      `json:"scheduledDate"`
}
