package model

import (
	"loanManagement/constant"

	"github.com/google/uuid"
)

type Payment struct {
	Base
	LoanID   uuid.UUID         `json:"loanID" gorm:"column:loan_id;type:uuid;"`
	UserID   uuid.UUID         `json:"userId" gorm:"column:user_id;type:uuid;"`
	Amount   int64             `json:"amount" gorm:"column:amount"`
	Currency constant.Currency `json:"currency" gorm:"column:currency;default:'INR'"`
	Loan     Loan              `json:"loan" gorm:"foreignKey:loan_id"`
	User     User              `json:"user" gorm:"foreignKey:user_id"`
}

func (Payment) TableName() string {
	return "payment"
}
