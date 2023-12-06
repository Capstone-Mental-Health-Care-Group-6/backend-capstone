package service

import (
	root "FinalProject/features/chat_messages"
	"FinalProject/features/chat_messages/dto"
	"time"

	"github.com/labstack/echo/v4"
)

type MessageService struct {
	data root.MessageDataInterface
}

func New(data root.MessageDataInterface) root.MessageServiceInterface {
	return &MessageService{
		data: data,
	}
}

func (s *MessageService) GetMessages(ctx echo.Context, chat int) []*dto.Response {
	// TODO: implement jwt authorization
	query := ctx.QueryParams()
	responses := []*dto.Response{}
	for _, data := range s.data.Get(chat, query) {
		response := &dto.Response{
			ID: data.ID,
			Sender: &dto.User{
				ID:    int(data.User.ID),
				Name:  data.User.Name,
				Email: data.User.Email,
			},
			Text:      data.Text,
			Blob:      data.Blob,
			Timestamp: data.Timestamp,
		}
		responses = append(responses, response)
	}
	return responses
}

func (s *MessageService) GetMessage(ctx echo.Context, chat int, message int) *dto.Response {
	// TODO: implement jwt authorization
	data := s.data.Find(chat, message)
	if data != nil {
		response := dto.Response{
			ID: data.ID,
			Sender: &dto.User{
				ID:    int(data.User.ID),
				Name:  data.User.Name,
				Email: data.User.Email,
			},
			Text:      data.Text,
			Blob:      data.Blob,
			Timestamp: data.Timestamp,
		}
		return &response
	}
	return nil
}

func (s *MessageService) CreateMessage(ctx echo.Context, chat int, request *dto.Request) *dto.Response {
	// TODO: implement jwt authorization
	data := &root.Message{
		ChatID:    chat,
		UserID:    request.Sender,
		Text:      request.Text,
		Blob:      request.Blob,
		Timestamp: time.Now(),
	}
	if result := s.data.Create(data); result != nil {
		return &dto.Response{
			ID: result.ID,
			Sender: &dto.User{
				ID:    int(result.User.ID),
				Name:  result.User.Name,
				Email: result.User.Email,
			},
			Text:      result.Text,
			Blob:      result.Blob,
			Timestamp: result.Timestamp,
		}
	}
	return nil
}

func (s *MessageService) UpdateMessage(ctx echo.Context, chat int, message int, request *dto.Request) *dto.Response {
	// TODO: implement jwt authorization
	data := &root.Message{
		ID:     message,
		ChatID: chat,
		UserID: request.Sender,
		Text:   request.Text,
		Blob:   request.Blob,
	}
	if result := s.data.Update(data); result != nil {
		return &dto.Response{
			ID: result.ID,
			Sender: &dto.User{
				ID:    int(result.User.ID),
				Name:  result.User.Name,
				Email: result.User.Email,
			},
			Text:      result.Text,
			Blob:      result.Blob,
			Timestamp: result.Timestamp,
		}
	}
	return nil
}

func (s *MessageService) DeleteMessage(ctx echo.Context, chat int, message int) bool {
	// TODO: implement jwt authorization
	return s.data.Delete(chat, message)
}
