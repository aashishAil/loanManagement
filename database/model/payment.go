package model

import "loanManagement/constant"

type Payment struct {
	Base
	LoanID   string            `json:"loanID" gorm:"column:loan_id"`
	Amount   int64             `json:"amount" gorm:"column:amount"`
	Currency constant.Currency `json:"currency" gorm:"column:currency;default:'INR'"`
	Loan     Loan              `json:"loan" gorm:"foreignKey:loan_id"`
}

func (Payment) TableName() string {
	return "payment"
}
