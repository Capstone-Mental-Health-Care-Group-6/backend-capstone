package handler

import "time"

type InputResponse struct {
	TransactionID uint      `json:"transaction_id" form:"transaction_id"`
	Date          time.Time `json:"date" form:"date"`
	Time          time.Time `json:"time" form:"time"`
	Duration      uint      `json:"duration" form:"duration"`
	Status        string    `json:"status" form:"status"`
}
