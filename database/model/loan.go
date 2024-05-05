package model

import (
	"time"

	"loanManagement/constant"

	"github.com/google/uuid"
)

type Loan struct {
	BaseWithUpdatedAt
	UserID            uuid.UUID           `json:"userId" gorm:"column:user_id;type:uuid;"`
	DisbursalAmount   int64               `json:"disbursalAmount" gorm:"column:disbursal_amount"`
	PendingAmount     int64               `json:"pendingAmount" gorm:"column:pending_amount"`
	WeeklyInstallment int64               `json:"weeklyInstallment" gorm:"column:weekly_installment"`
	Currency          constant.Currency   `json:"currency" gorm:"column:currency;default:'INR'"`
	Term              int64               `json:"term" gorm:"column:term"`
	Status            constant.LoanStatus `json:"status" gorm:"column:status;default:'PENDING'"`
	DisbursalDate     time.Time           `json:"disbursalDate" gorm:"column:disbursal_date;default:(now() at time zone 'utc')"`
	User              User                `json:"user" gorm:"foreignKey:user_id"`
}
