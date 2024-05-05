package model

import (
	"time"

	"loanManagement/constant"
)

type ScheduleRepayment struct {
	BaseWithUpdatedAt
	LoanID          string                         `json:"loanID" gorm:"column:loan_id"`
	ScheduledAmount int64                          `json:"scheduledAmount" gorm:"column:scheduled_amount"`
	PendingAmount   int64                          `json:"pendingAmount" gorm:"column:pending_amount"`
	Currency        constant.Currency              `json:"currency" gorm:"column:currency;default:'INR'"`
	Status          constant.SchedulePaymentStatus `json:"status" gorm:"column:status;default:'PENDING'"`
	ScheduledDate   time.Time                      `json:"scheduledDate" gorm:"column:scheduled_date"`
	Loan            Loan                           `json:"loan" gorm:"foreignKey:loan_id"`
}

func (ScheduleRepayment) TableName() string {
	return "schedule_repayment"
}
