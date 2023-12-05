package service

import (
	root "FinalProject/features/chats"
	"FinalProject/features/chats/dto"
	"FinalProject/utils/websocket"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

type ChatService struct {
	sock *websocket.Server
	data root.ChatDataInterface
}

func New(data root.ChatDataInterface, sock *websocket.Server) root.ChatServiceInterface {
	return &ChatService{
		sock: sock,
		data: data,
	}
}

func (s *ChatService) SocketEstablish(ctx echo.Context, user int) *websocket.Client {
	// TODO: implement jwt authorization
	client := s.sock.CreateClient(ctx, user)
	for _, data := range s.data.Get(user, nil) {
		if s.sock.FindRoom(data.ID) == nil {
			s.sock.CreateRoom(data.ID, user)
		} else {
			s.sock.JoinRoom(data.ID, user)
		}
	}
	ctx.Set(fmt.Sprintf("ws#%d", user), true)
	return client
}

func (s *ChatService) GetChats(ctx echo.Context, user int) []*dto.GetChatResponse {
	// TODO: implement jwt authorization
	if s.sock.FindClient(user) == nil {
		ctx.Set("websocket.connection.error", true)
		return nil
	}
	query := ctx.QueryParams()
	responses := []*dto.GetChatResponse{}
	for _, data := range s.data.Get(user, query) {
		if s.sock.FindRoom(data.ID) == nil {
			s.sock.CreateRoom(data.ID, user)
		} else {
			s.sock.JoinRoom(data.ID, user)
		}
		response := &dto.GetChatResponse{
			ID: data.ID,
			Patient: &dto.User{
				ID:    int(data.Patient.ID),
				Name:  data.Patient.Name,
				Email: data.Patient.Email,
			},
			Doctor: &dto.User{
				ID:    int(data.Doctor.ID),
				Name:  data.Doctor.Name,
				Email: data.Doctor.Email,
			},
			LastMessage:       data.LastMessage,
			LastMessageTime:   data.LastMessageTime,
			LastMessageSentBy: data.LastMessageSentByID,
			LastMessageSeenBy: data.LastMessageSeenByID,
		}
		responses = append(responses, response)
	}
	return responses
}

func (s *ChatService) CreateChat(ctx echo.Context, request *dto.CreateChatRequest) *dto.CreateChatResponse {
	// TODO: implement jwt authorization
	// TODO: check websocket connection
	data := &root.Chat{
		PatientUserID: request.Patient,
		DoctorUserID:  request.Doctor,
	}
	if result := s.data.Create(data); result != nil {
		var room *websocket.Room
		if room = s.sock.FindRoom(result.ID); room == nil {
			s.sock.CreateRoom(result.ID, int(result.Patient.ID), int(result.Doctor.ID))
		}
		return &dto.CreateChatResponse{
			ID: result.ID,
			Patient: &dto.User{
				ID:    int(result.Patient.ID),
				Name:  result.Patient.Name,
				Email: result.Patient.Email,
			},
			Doctor: &dto.User{
				ID:    int(result.Doctor.ID),
				Name:  result.Doctor.Name,
				Email: result.Doctor.Email,
			},
		}
	}
	return nil
}

func (s *ChatService) UpdateChat(ctx echo.Context, chat int, request *dto.UpdateChatRequest) *dto.UpdateChatResponse {
	// TODO: implement jwt authorization
	// TODO: check websocket connection
	data := &root.Chat{
		ID:                  chat,
		LastMessage:         request.Message,
		LastMessageTime:     func() *time.Time { t := time.Now(); return &t }(),
		LastMessageSentByID: request.SentBy,
		LastMessageSeenByID: request.SeenBy,
	}
	if result := s.data.Update(data); result != nil {
		return &dto.UpdateChatResponse{
			ID:                result.ID,
			LastMessage:       result.LastMessage,
			LastMessageTime:   result.LastMessageTime,
			LastMessageSentBy: result.LastMessageSentByID,
			LastMessageSeenBy: result.LastMessageSeenByID,
		}
	}
	return nil
}

func (s *ChatService) DeleteChat(ctx echo.Context, chat int) bool {
	// TODO: implement jwt authorization
	// TODO: check websocket connection
	if s.sock.FindRoom(chat) != nil {
		s.sock.DeleteRoom(chat)
	}
	return s.data.Delete(chat)
}
