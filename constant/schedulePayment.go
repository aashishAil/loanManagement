package constant

type ScheduleRepaymentStatus string

const (
	ScheduleRepaymentStatusPending  ScheduleRepaymentStatus = "PENDING"
	ScheduleRepaymentStatusApproved ScheduleRepaymentStatus = "APPROVED"
	ScheduleRepaymentStatusRejected ScheduleRepaymentStatus = "REJECTED"
	ScheduleRepaymentStatusPaid     ScheduleRepaymentStatus = "PAID"
)
