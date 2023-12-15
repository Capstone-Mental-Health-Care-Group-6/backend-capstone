package handler

import "time"

type InputRequest struct {
	TransactionID uint      `json:"transaction_id" form:"transaction_id"`
	Date          time.Time `json:"date" form:"date"`
	Time          time.Time `json:"time" form:"time"`
	Duration      uint      `json:"duration" form:"duration"`
}

type InputRequestUpdate struct {
	TransactionID uint      `json:"transaction_id" form:"transaction_id"`
	Date          time.Time `json:"date" form:"date"`
	Time          time.Time `json:"time" form:"time"`
	Duration      uint      `json:"duration" form:"duration"`
	Status        string    `json:"status" form:"status"`
}

type RequestStatusUpdate struct {
	Alasan string `json:"alasan" form:"alasan"`
}
