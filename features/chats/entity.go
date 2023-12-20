package chats

import (
	"FinalProject/features/chats/dto"
	"net/url"

	"github.com/labstack/echo/v4"
)

type ChatHandlerInterface interface {
	Establish() echo.HandlerFunc
	Index() echo.HandlerFunc
	Store() echo.HandlerFunc
	Edit() echo.HandlerFunc
	Destroy() echo.HandlerFunc
}

type ChatServiceInterface interface {
	SocketEstablish(ctx echo.Context, user int, role string)
	GetChats(ctx echo.Context) []*dto.GetChatResponse
	CreateChat(ctx echo.Context, request *dto.CreateChatRequest) *dto.CreateChatResponse
	UpdateChat(ctx echo.Context, chat int, request *dto.UpdateChatRequest) *dto.UpdateChatResponse
	DeleteChat(ctx echo.Context, chat int) bool
}

type ChatDataInterface interface {
	Get(user int, role string, query url.Values) []*Chat
	Find(chat int) *Chat
	Create(data *Chat) *Chat
	Update(data *Chat) *Chat
	Delete(chat int) bool
}
