package chatbot

import (
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chatbot struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID       uint               `json:"user_id"`
	Prompt       string             `json:"prompt"`
	ResultPrompt string             `json:"result_prompt"`
	Date         time.Time          `json:"date"`
}

type ChatbotHandlerInterface interface {
	GetAllChatBot() echo.HandlerFunc
	CreateChatBot() echo.HandlerFunc
}

type ChatbotServiceInterface interface {
	GetAllChatBot(user_id int) ([]Chatbot, error)
	InsertChatBot(input Chatbot) (Chatbot, error)
}

type ChatbotDataInterface interface {
	GetAllChatBot(user_id int) ([]Chatbot, error)
	InsertChatBot(input Chatbot) (Chatbot, error)
}
