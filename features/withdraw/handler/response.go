package handler

type WithdrawResponse struct {
	DoctorID      uint   `json:"doctor_id"`
	BalanceReq    uint   `json:"balance_req"`
	BalanceBefore uint   `json:"balance_before"`
	BalanceAfter  uint   `json:"balance_after"`
	PaymentMethod string `json:"payment_method"`
	PaymentNumber string `json:"payment_number"`
	PaymentName   string `json:"payment_name"`
}
