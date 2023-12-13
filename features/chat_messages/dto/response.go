package dto

import "time"

type User struct {
	ID    int    `json:"user_id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
	// Phone string `json:"phone,omitempty"`	# reserved
	// Photo string `json:"photo"`				# reserved
}

type Response struct {
	ID        int       `json:"message_id"`
	Sender    *User     `json:"sender"`
	Text      string    `json:"text"`
	Blob      string    `json:"blob"`
	Timestamp time.Time `json:"timestamp"`
}
