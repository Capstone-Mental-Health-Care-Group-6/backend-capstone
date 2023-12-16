package dto

import "time"

type User struct {
	ID    int    `json:"user_id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
}

type GetChatResponse struct {
	ID                int        `json:"chat_id"`
	Patient           *User      `json:"patient"`
	Doctor            *User      `json:"doctor"`
	LastMessage       string     `json:"last_message"`
	LastMessageTime   *time.Time `json:"last_message_time"`
	LastMessageSentBy *int       `json:"last_message_sent_by"`
	LastMessageSeenBy *int       `json:"last_message_seen_by"`
}

type CreateChatResponse struct {
	ID      int   `json:"chat_id"`
	Patient *User `json:"patient"`
	Doctor  *User `json:"doctor"`
}

type UpdateChatResponse struct {
	ID                int        `json:"chat_id"`
	LastMessage       string     `json:"last_message"`
	LastMessageTime   *time.Time `json:"last_message_time"`
	LastMessageSentBy *int       `json:"last_message_sent_by"`
	LastMessageSeenBy *int       `json:"last_message_seen_by"`
}
