package service

import (
	"FinalProject/features/chatbotcs"
)

type ChatbotCsService struct {
	channels []chan string
	data     chatbotcs.ChatbotCsDataInterface
}

func New(data chatbotcs.ChatbotCsDataInterface) chatbotcs.ChatbotCsServiceInterface {
	return &ChatbotCsService{
		channels: make([]chan string, 0),
		data:     data,
	}
}

func (s *ChatbotCsService) CreateMsg(message string) {
	for _, channel := range s.channels {
		channel <- message
	}
}

func (s *ChatbotCsService) JoinGroup() chan string {
	channel := make(chan string)
	s.channels = append(s.channels, channel)

	return channel
}

func (s *ChatbotCsService) LeaveGroup(channel chan string) {
	close(channel)
	for i, ch := range s.channels {
		if ch == channel {
			s.channels = append(s.channels[:i], s.channels[i+1:]...)
			break
		}
	}
}

func (s *ChatbotCsService) GetAnswer(question string) string {
	return s.data.GetAnswer(question)
}
