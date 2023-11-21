package handler

type InputRequest struct {
	DoctorID      uint   `json:"doctor_id"`
	BalanceReq    uint   `json:"balance_req"`
	PaymentMethod string `json:"payment_method"`
	PaymentNumber string `json:"payment_number"`
}

type UpdateRequest struct {
	ConfirmByID uint   `json:"confirm_by_id"`
	Status      string `json:"status"`
}
