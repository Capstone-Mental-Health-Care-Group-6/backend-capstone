package data

import "time"

type Chatbot struct {
	ID           string    `json:"id"`
	UserID       uint      `json:"user_id"`
	Prompt       string    `json:"prompt"`
	ResultPrompt string    `json:"result_prompt"`
	Date         time.Time `json:"date"`
}
