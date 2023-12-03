package service

import (
	"FinalProject/features/chatbot"
	"FinalProject/utils/openai"
)

type ChatbotService struct {
	data   chatbot.ChatbotDataInterface
	openai openai.OpenAIInterface
}

func New(data chatbot.ChatbotDataInterface, openai openai.OpenAIInterface) chatbot.ChatbotServiceInterface {
	return &ChatbotService{
		data:   data,
		openai: openai,
	}
}
