package handler

type ChargeTransactionResponse struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
}

type ManualTransactionResponse struct {
	PriceResult   uint   `json:"price_result"`
	UserID        uint   `json:"user_id"`
	MidtransID    string `json:"transaction_id"`
	PaymentStatus uint   `json:"payment_status"`
	PaymentProof  string `json:"payment_proof"`
	PaymentType   string `json:"payment_type"`
}
