package model

import (
	"time"

	"loanManagement/constant"
	routerModel "loanManagement/router/model"

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

func (*Loan) TableName() string {
	return "loan"
}

func (model *Loan) TransformForRouter(repayments []*ScheduledRepayment) routerModel.UserLoan {
	loanI := routerModel.UserLoan{
		ID:              model.ID,
		UserID:          model.UserID,
		DisbursalAmount: float64(model.DisbursalAmount) / constant.MinCurrencyConversionFactor,
		PendingAmount:   float64(model.PendingAmount) / constant.MinCurrencyConversionFactor,
		Currency:        model.Currency,
		Term:            model.Term,
		Status:          model.Status,
		DisbursalDate:   model.DisbursalDate,
	}

	if len(repayments) > 0 {
		repaymentArr := make([]routerModel.UserScheduledRepayment, len(repayments))
		for i := range repayments {
			repayment := repayments[i]
			repaymentArr[i] = repayment.TransformForRouter()
		}

		loanI.ScheduledRepayments = &repaymentArr
	}

	return loanI
}
