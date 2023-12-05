package chat_messages

import (
	"FinalProject/features/chat_messages/dto"
	"net/url"

	"github.com/labstack/echo/v4"
)

type MessageHandlerInterface interface {
	Index() echo.HandlerFunc
	Observe() echo.HandlerFunc
	Store() echo.HandlerFunc
	Edit() echo.HandlerFunc
	Destroy() echo.HandlerFunc
}

type MessageServiceInterface interface {
	GetMessages(ctx echo.Context, chat int) []*dto.Response
	GetMessage(ctx echo.Context, chat int, message int) *dto.Response
	CreateMessage(ctx echo.Context, chat int, request *dto.Request) *dto.Response
	UpdateMessage(ctx echo.Context, chat int, message int, request *dto.Request) *dto.Response
	DeleteMessage(ctx echo.Context, chat int, message int) bool
}

type MessageDataInterface interface {
	Get(chat int, query url.Values) []*Message
	Find(chat int, message int) *Message
	Create(data *Message) *Message
	Update(data *Message) *Message
	Delete(chat int, message int) bool
}
