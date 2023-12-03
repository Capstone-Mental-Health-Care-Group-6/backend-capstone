package chatbot

type Chatbot struct {
	UserID uint   `json:"user_id"`
	Prompt string `json:"prompt"`
	Answer string `json:"answer"`
}

type ChatbotHandlerInterface interface {
}

type ChatbotServiceInterface interface {
}

type ChatbotDataInterface interface {
}
