package constant

type LoanStatus string

const (
	LoanStatusPending  LoanStatus = "PENDING"
	LoanStatusApproved LoanStatus = "APPROVED"
	LoanStatusRejected LoanStatus = "REJECTED"
	LoanStatusPaid     LoanStatus = "PAID"

	LoanCreationDisbursalTimeGap = 7  // in days
	LoanRepaymentTimeGap         = 7  // in days
	MinLoanRepaymentTenure       = 4  // in weeks
	MaxLoanRepaymentTenure       = 52 // in weeks
	MinDisbursalAmount           = 1000
	MaxDisbursalAmount           = 10_00_00_000
)
