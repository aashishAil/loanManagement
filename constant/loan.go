package constant

type LoanStatus string

const (
	LoanStatusPending  LoanStatus = "PENDING"
	LoanStatusApproved LoanStatus = "APPROVED"
	LoanStatusRejected LoanStatus = "REJECTED"
	LoanStatusPaid     LoanStatus = "PAID"
)
