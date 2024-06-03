package constant

type LoanStatus string

const (
	LoanStatusPending  LoanStatus = "PENDING"
	LoanStatusApproved LoanStatus = "APPROVED"
	LoanStatusRejected LoanStatus = "REJECTED"
	LoanStatusPaid     LoanStatus = "PAID"

	LoanCreationDisbursalTimeGap = 7  // in days
	LoanRepaymentTimeGap         = 7  // in days
	MaxLoanRepaymentTenure       = 52 // in weeks
)
