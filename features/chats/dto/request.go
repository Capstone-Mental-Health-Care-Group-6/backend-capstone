package dto

type CreateChatRequest struct {
	Patient int `json:"patient"`
	Doctor  int `json:"doctor"`
}

type UpdateChatRequest struct {
	Message string `json:"message,omitempty"`
	SentBy  *int   `json:"sent_by,omitempty"`
	SeenBy  *int   `json:"seen_by,omitempty"`
}
