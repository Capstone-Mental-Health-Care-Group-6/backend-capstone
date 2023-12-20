package service

import (
	root "FinalProject/features/chats"
	"FinalProject/features/chats/dto"
	"FinalProject/features/patients"
	"FinalProject/features/users"
	"FinalProject/helper"
	"FinalProject/utils/websocket"
	"fmt"
	"strings"
	"time"

	"github.com/fatih/structs"
	"github.com/labstack/echo/v4"
)

type ChatService struct {
	jwt     helper.JWTInterface
	socket  *websocket.Server
	chat    root.ChatDataInterface
	patient patients.PatientDataInterface
	doctor  users.UserDataInterface
}

func New(
	dataChat root.ChatDataInterface,
	dataPatient patients.PatientDataInterface,
	dataDoctor users.UserDataInterface,
	socket *websocket.Server,
	jwt helper.JWTInterface,
) root.ChatServiceInterface {
	return &ChatService{
		jwt:     jwt,
		socket:  socket,
		chat:    dataChat,
		patient: dataPatient,
		doctor:  dataDoctor,
	}
}

func (s *ChatService) SocketEstablish(ctx echo.Context, user int, role string) {
	res, err := func() (res any, err error) {
		switch role {
		case "patient":
			res, err = s.patient.GetByID(user)
		case "doctor":
			res, err = s.doctor.GetByID(user)
		}
		return
	}()
	if err != nil || structs.IsZero(res) {
		ctx.Set("ws.client.error", "user not found")
		return
	}
	sign := s.socket.CreateClient(ctx, user, role)
	ref := fmt.Sprintf("%s@%d", role, user)
	for _, data := range s.chat.Get(user, role, nil) {
		if s.socket.FindRoom(data.ID) == nil {
			s.socket.CreateRoom(data.ID, ref)
		} else {
			s.socket.JoinRoom(data.ID, ref)
		}
	}
	ctx.Set("ws.connect", sign)
}

func (s *ChatService) GetChats(ctx echo.Context) []*dto.GetChatResponse {
	user, err := s.jwt.GetID(ctx)
	if err != nil {
		ctx.Set("jwt.token.error", true)
		return nil
	}
	role := strings.ToLower(s.jwt.CheckRole(ctx).(string))
	ref := fmt.Sprintf("%s@%d", role, user)
	if s.socket.FindClient(ref) == nil {
		ctx.Set("ws.connect.error", true)
		return nil
	}
	query := ctx.QueryParams()
	responses := []*dto.GetChatResponse{}
	for _, data := range s.chat.Get(int(user), role, query) {
		if s.socket.FindRoom(data.ID) == nil {
			s.socket.CreateRoom(data.ID, ref)
		} else {
			s.socket.JoinRoom(data.ID, ref)
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
	var res, err any
	res, err = s.patient.GetByID(request.Patient)
	if err != nil || structs.IsZero(res) {
		ctx.Set("ws.client.error", fmt.Sprintf("patient@%d not found", request.Patient))
		return nil
	}
	res, err = s.doctor.GetByID(request.Doctor)
	if err != nil || structs.IsZero(res) {
		ctx.Set("ws.client.error", fmt.Sprintf("doctor@%d not found", request.Doctor))
		return nil
	}
	// user, err := s.jwt.GetID(ctx)
	// if err != nil {
	// 	ctx.Set("jwt.token.error", true)
	// 	return nil
	// }
	// role := strings.ToLower(s.jwt.CheckRole(ctx).(string))
	// ref := fmt.Sprintf("%s@%d", role, user)
	// if s.socket.FindClient(ref) == nil {
	// 	ctx.Set("ws.connect.error", true)
	// 	return nil
	// }
	patient_ref := fmt.Sprintf("patient@%d", request.Patient)
	if s.socket.FindClient(patient_ref) == nil {
		ctx.Set("ws.connect.error", fmt.Sprintf("%s not connected yet", patient_ref))
		return nil
	}
	doctor_ref := fmt.Sprintf("doctor@%d", request.Doctor)
	if s.socket.FindClient(doctor_ref) == nil {
		ctx.Set("ws.connect.error", fmt.Sprintf("%s not connected yet", doctor_ref))
		return nil
	}
	data := &root.Chat{
		PatientUserID: request.Patient,
		DoctorUserID:  request.Doctor,
	}
	if result := s.chat.Create(data); result != nil {
		if room := s.socket.FindRoom(result.ID); room == nil {
			s.socket.CreateRoom(result.ID, patient_ref, doctor_ref)
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
	if s.socket.FindClient(ref) == nil {
		ctx.Set("ws.connect.error", fmt.Sprintf("%s not connected yet", ref))
		return nil
	}
	data := &root.Chat{
		ID:                  chat,
		LastMessage:         request.Message,
		LastMessageTime:     func() *time.Time { t := time.Now(); return &t }(),
		LastMessageSentByID: request.SentBy,
		LastMessageSeenByID: request.SeenBy,
	}
	if result := s.chat.Update(data); result != nil {
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
	if s.socket.FindClient(ref) == nil {
		ctx.Set("ws.connect.error", fmt.Sprintf("%s not connected yet", ref))
		return false
	}
	if s.socket.FindRoom(chat) != nil {
		s.socket.DeleteRoom(chat)
	}
	return s.chat.Delete(chat)
}
