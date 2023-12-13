package service

import (
	"FinalProject/features/chatbot"
	"FinalProject/utils/openai"
	"errors"
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

func (s *ChatbotService) GetAllChatBot(user_id int) ([]chatbot.Chatbot, error) {
	data, err := s.data.GetAllChatBot(user_id)
	if err != nil {
		return data, errors.New("Get All Process Failed")
	}
	return data, nil
}

func (s *ChatbotService) InsertChatBot(input chatbot.Chatbot) (chatbot.Chatbot, error) {

	openaiResult, err := s.openai.GenerateText(input.Prompt)
	if err != nil {
		return chatbot.Chatbot{}, errors.New("Generate Text Error")
	}

	input.ResultPrompt = openaiResult

	data, err := s.data.InsertChatBot(input)
	if err != nil {
		return data, errors.New("Insert Process Failed")
	}
	return data, nil
}
