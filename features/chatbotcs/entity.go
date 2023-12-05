package chatbotcs

import "github.com/labstack/echo/v4"

type ChatbotCs struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

type ChatbotCsHandlerInterface interface {
	ChatBotCs() echo.HandlerFunc
	CreateMessage() echo.HandlerFunc
}

type ChatbotCsServiceInterface interface {
	CreateMsg(message string)
	JoinGroup() chan string
	LeaveGroup(channel chan string)
	GetAnswer(question string) string
}

type ChatbotCsDataInterface interface {
	GetAnswer(question string) string
}
