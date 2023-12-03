package data

import (
	"FinalProject/features/chatbot"
)

type ChatbotData struct {
}

func New() chatbot.ChatbotDataInterface {
	return &ChatbotData{}
}
