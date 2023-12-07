package service

import (
	"FinalProject/features/chatbotcs"
)

type ChatbotCsService struct {
	channels map[string]chan chatbotcs.ChatbotCs
	data     chatbotcs.ChatbotCsDataInterface
}

func New(data chatbotcs.ChatbotCsDataInterface) chatbotcs.ChatbotCsServiceInterface {
	return &ChatbotCsService{
		channels: make(map[string]chan chatbotcs.ChatbotCs),
		data:     data,
	}
}

func (s *ChatbotCsService) CreateMsg(ip string, message chatbotcs.ChatbotCs) {
	if channel, ok := s.channels[ip]; ok {
		channel <- message
	}
}

func (s *ChatbotCsService) JoinGroup(ip string) chan chatbotcs.ChatbotCs {
	channel := make(chan chatbotcs.ChatbotCs)
	s.channels[ip] = channel
	return channel
}

func (s *ChatbotCsService) LeaveGroup(ip string) {
	if channel, ok := s.channels[ip]; ok {
		close(channel)
		delete(s.channels, ip)
	}
}

func (s *ChatbotCsService) GetAnswer(question string) string {
	return s.data.GetAnswer(question)
}
