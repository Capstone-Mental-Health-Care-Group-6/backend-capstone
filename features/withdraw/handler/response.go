package handler

type WithdrawResponse struct {
	DoctorID      uint
	BalanceReq    uint
	BalanceBefore uint
	BalanceAfter  uint
	PaymentMethod string
	PaymentNumber string
	PaymentName   string
}
