package handler

type InputRequest struct {
	BalanceReq    uint   `json:"balance_req" validate:"required"`
	PaymentMethod string `json:"payment_method" validate:"required"`
	PaymentNumber string `json:"payment_number" validate:"required"`
	PaymentName   string `json:"payment_name" validate:"required"`
}

type UpdateRequest struct {
	ConfirmByID uint   `json:"confirm_by_id"`
	Status      string `json:"status"`
}
