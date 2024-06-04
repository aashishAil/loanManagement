package model

import (
	routerModel "loanManagement/router/model"
	"time"

	"loanManagement/constant"

	"github.com/google/uuid"
)

type ScheduledRepayment struct {
	BaseWithUpdatedAt
	LoanID          uuid.UUID                        `json:"loanID" gorm:"column:loan_id;type:uuid;"`
	ScheduledAmount int64                            `json:"scheduledAmount" gorm:"column:scheduled_amount"`
	PendingAmount   int64                            `json:"pendingAmount" gorm:"column:pending_amount"`
	Currency        constant.Currency                `json:"currency" gorm:"column:currency;default:'INR'"`
	Status          constant.ScheduleRepaymentStatus `json:"status" gorm:"column:status;default:'PENDING'"`
	ScheduledDate   time.Time                        `json:"scheduledDate" gorm:"column:scheduled_date"`
	Loan            Loan                             `json:"loan" gorm:"foreignKey:loan_id"`
}

func (*ScheduledRepayment) TableName() string {
	return "scheduled_repayment"
}

func (model *ScheduledRepayment) TransformForRouter() routerModel.UserScheduledRepayment {
	return routerModel.UserScheduledRepayment{
		ID:              model.ID,
		LoanID:          model.LoanID,
		ScheduledAmount: float64(model.ScheduledAmount) / constant.MinCurrencyConversionFactor,
		PendingAmount:   float64(model.PendingAmount) / constant.MinCurrencyConversionFactor,
		Currency:        model.Currency,
		Status:          model.Status,
		ScheduledDate:   model.ScheduledDate,
	}
}
