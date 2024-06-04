package constant

import "github.com/pkg/errors"

var (
	InvalidDisbursalDateGapError = errors.New("disbursal date should be at least 7 days from today")
	MinRepaymentTenureError      = errors.New("loan repayment tenure cannot be less than 4 weeks")
	MaxRepaymentTenureError      = errors.New("loan repayment tenure cannot be more than 52 weeks")
	InvalidMinAmountError        = errors.New("amount should be at least 1000")
	InvalidMaxAmountError        = errors.New("amount should be at most 10,00,00,000")
	InvalidCurrencyError         = errors.New("invalid currency provided")
	LoanInTerminalStatusError    = errors.New("loan is in terminal status")
)
