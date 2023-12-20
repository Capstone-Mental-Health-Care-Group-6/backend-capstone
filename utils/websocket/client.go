package websocket

import (
	"encoding/base64"
	"fmt"
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
	sign    string
}

func NewClient(context echo.Context, server *Server, user int, role string) (string, *Client) {
	var (
		format = fmt.Sprintf("%s@%d", role, user)
		sign   = base64.RawStdEncoding.EncodeToString([]byte(format))
	)
	return format, &Client{
		server:  server,
		handler: NewProtocol().Switch(context),
		rooms:   make(map[int]*Room),
		message: make(chan *packet.Message),
		sign:    sign,
	}
}

func (c *Client) disconnect() {
	c.handler.Close()
	c.server.DeleteClient(c.sign)
	close(c.message)
	logrus.Infof("[ws.establish]: client@%s disconnected", c.sign)
}

func (c *Client) Send() {
	var (
		duration time.Duration = time.Second * 10
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
		if c.rooms[packet.Room] != nil {
			c.rooms[packet.Room].message <- packet
		} else {
			fmt.Println("Chat Room", packet.Room, "belum ada")
		}
	}
}

func (c *Client) Recv() {
	var (
		duration time.Duration = time.Second * 9
		ticker   *time.Ticker  = time.NewTicker(duration)
		err      error
	)
	defer c.handler.Close()
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
