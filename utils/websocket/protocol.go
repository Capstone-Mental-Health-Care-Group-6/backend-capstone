package websocket

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Protocol struct {
	upgrader *websocket.Upgrader
	dialer   *websocket.Dialer
}

func NewProtocol() *Protocol {
	return &Protocol{
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  4096,
			WriteBufferSize: 4096,
		},
		dialer: websocket.DefaultDialer,
	}
}

func (proto *Protocol) Switch(ctx echo.Context) *websocket.Conn {
	response, request := ctx.Response(), ctx.Request()
	header := http.Header{}
	handler, err := proto.upgrader.Upgrade(response, request, header)
	if err != nil {
		logrus.Error("[protocol.switch]: ", err.Error())
		return nil
	}
	return handler
}
