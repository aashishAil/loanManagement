package model

import "loanManagement/constant"

type GetAdminLoansOutput struct {
	Loans []UserLoan `json:"loans"`
}

type UpdateAdminLoanInput struct {
	Status constant.LoanStatus `json:"status"`
}
