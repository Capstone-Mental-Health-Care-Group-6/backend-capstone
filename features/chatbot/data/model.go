package data

type Chatbot struct {
	UserID uint   `json:"user_id"`
	Prompt string `json:"prompt"`
	Answer string `json:"answer"`
}
