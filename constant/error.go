package constant

import "github.com/pkg/errors"

var (
	InvalidDisbursalDateGapError = errors.New("disbursal date should be at least 7 days from today")
	MaxRepaymentTenureError      = errors.New("loan repayment tenure cannot be more than 52 weeks")
)
