package service

import (
	root "FinalProject/features/chats"
	"FinalProject/features/chats/dto"
	"FinalProject/helper"
	"FinalProject/utils/websocket"
	"fmt"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type ChatService struct {
	jwt  helper.JWTInterface
	sock *websocket.Server
	data root.ChatDataInterface
}

func New(data root.ChatDataInterface, sock *websocket.Server, jwt helper.JWTInterface) root.ChatServiceInterface {
	return &ChatService{
		jwt:  jwt,
		sock: sock,
		data: data,
	}
}

func (s *ChatService) SocketEstablish(ctx echo.Context, user int, role string) {
	ref := s.sock.CreateClient(ctx, user, role)
	for _, data := range s.data.Get(user, role, nil) {
		if s.sock.FindRoom(data.ID) == nil {
			s.sock.CreateRoom(data.ID, ref)
		} else {
			s.sock.JoinRoom(data.ID, ref)
		}
	}
	ctx.Set("ws.connect", ref)
}

func (s *ChatService) GetChats(ctx echo.Context, user int) []*dto.GetChatResponse {
	role := strings.ToLower(s.jwt.CheckRole(ctx).(string))
	ref := fmt.Sprintf("%s@%d", role, user)
	if s.sock.FindClient(ref) == nil {
		ctx.Set("ws.connect.error", true)
		return nil
	}
	query := ctx.QueryParams()
	responses := []*dto.GetChatResponse{}
	for _, data := range s.data.Get(user, role, query) {
		if s.sock.FindRoom(data.ID) == nil {
			s.sock.CreateRoom(data.ID, ref)
		} else {
			s.sock.JoinRoom(data.ID, ref)
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
	user, err := s.jwt.GetID(ctx)
	if err != nil {
		ctx.Set("jwt.token.error", true)
		return nil
	}
	role := strings.ToLower(s.jwt.CheckRole(ctx).(string))
	ref := fmt.Sprintf("%s@%d", role, user)
	if s.sock.FindClient(ref) == nil {
		ctx.Set("ws.connect.error", true)
		return nil
	}
	data := &root.Chat{
		PatientUserID: request.Patient,
		DoctorUserID:  request.Doctor,
	}
	patient_ref := fmt.Sprintf("patient@%d", request.Patient)
	if s.sock.FindClient(patient_ref) == nil {
		ctx.Set("ws.connect.error", true)
		return nil
	}
	doctor_ref := fmt.Sprintf("doctor@%d", request.Doctor)
	if s.sock.FindClient(doctor_ref) == nil {
		ctx.Set("ws.connect.error", true)
		return nil
	}
	if result := s.data.Create(data); result != nil {
		if room := s.sock.FindRoom(result.ID); room == nil {
			s.sock.CreateRoom(result.ID, patient_ref, doctor_ref)
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
	user, err := s.jwt.GetID(ctx)
	if err != nil {
		ctx.Set("jwt.token.error", true)
		return nil
	}
	role := strings.ToLower(s.jwt.CheckRole(ctx).(string))
	ref := fmt.Sprintf("%s@%d", role, user)
	if s.sock.FindClient(ref) == nil {
		ctx.Set("ws.connect.error", true)
		return nil
	}
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
	user, err := s.jwt.GetID(ctx)
	if err != nil {
		ctx.Set("jwt.token.error", true)
		return false
	}
	role := strings.ToLower(s.jwt.CheckRole(ctx).(string))
	ref := fmt.Sprintf("%s@%d", role, user)
	if s.sock.FindClient(ref) == nil {
		ctx.Set("ws.connect.error", true)
		return false
	}
	if s.sock.FindRoom(chat) != nil {
		s.sock.DeleteRoom(chat)
	}
	return s.data.Delete(chat)
}
