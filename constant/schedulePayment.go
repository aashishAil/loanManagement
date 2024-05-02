package constant

type SchedulePaymentStatus string

const (
	SchedulePaymentStatusPending  SchedulePaymentStatus = "PENDING"
	SchedulePaymentStatusApproved SchedulePaymentStatus = "APPROVED"
	SchedulePaymentStatusPaid     SchedulePaymentStatus = "PAID"
)
