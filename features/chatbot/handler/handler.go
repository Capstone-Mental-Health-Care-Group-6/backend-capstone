package handler

import (
	"FinalProject/features/chatbot"
	"FinalProject/helper"
)

type ChatbotHandler struct {
	svc chatbot.ChatbotServiceInterface
	jwt helper.JWTInterface
}

func New(svc chatbot.ChatbotServiceInterface, jwt helper.JWTInterface) chatbot.ChatbotHandlerInterface {
	return &ChatbotHandler{
		svc: svc,
		jwt: jwt,
	}
}
