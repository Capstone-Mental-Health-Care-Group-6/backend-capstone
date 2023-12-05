package websocket

import (
	"time"

	"FinalProject/utils/websocket/packet"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Client struct {
	server  *Server
	handler *websocket.Conn
	rooms   map[int]*Room
	message chan *packet.Message
	sign    int
}

func NewClient(context echo.Context, server *Server, id int) *Client {
	return &Client{
		server:  server,
		handler: NewProtocol().Switch(context),
		rooms:   make(map[int]*Room),
		message: make(chan *packet.Message),
		sign:    id,
	}
}

func (c *Client) disconnect() {
	c.handler.Close()
	c.server.DeleteClient(c.sign)
}

func (c *Client) Send() {
	var (
		duration time.Duration = time.Second * 60
		packet   *packet.Message
		err      error
	)
	defer c.disconnect()
	c.handler.SetReadLimit(4096)
	c.handler.SetReadDeadline(time.Now().Add(duration))
	c.handler.SetPongHandler(func(string) error {
		c.handler.SetReadDeadline(time.Now().Add(duration))
		return nil
	})
	for {
		err = c.handler.ReadJSON(&packet)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, []int{
				websocket.CloseNormalClosure,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure,
			}...) {
				logrus.Error("[participant.send]: ", err.Error())
			}
			break
		}
		packet.Time = time.Now()
		c.rooms[packet.Room].message <- packet
	}
}

func (c *Client) Recv() {
	var (
		duration time.Duration = time.Second * 55
		ticker   *time.Ticker  = time.NewTicker(duration)
		err      error
	)
	defer c.disconnect()
	defer ticker.Stop()
	for {
		duration := time.Second * 15
		select {
		case message := <-c.message:
			c.handler.SetWriteDeadline(time.Now().Add(duration))
			err = c.handler.WriteJSON(message)
			if err != nil {
				logrus.Error("[participant.recv]: ", err.Error())
				return
			}
		case <-ticker.C:
			c.handler.SetWriteDeadline(time.Now().Add(duration))
			err = c.handler.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				return
			}
		}
	}
}
