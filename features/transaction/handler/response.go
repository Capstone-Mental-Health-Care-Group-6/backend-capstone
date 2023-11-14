package handler

type ChargeTransactionResponse struct {
	OrderID string `json:"order_id"`
	Status  string `json:"status"`
}
