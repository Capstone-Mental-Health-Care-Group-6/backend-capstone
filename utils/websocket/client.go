package websocket

import (
	"time"

	"FinalProject/utils/websocket/packet"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Participant struct {
	handler *websocket.Conn
	message chan *packet.Message
	room    *Room
	id      int
}

func NewParticipant(id int, room *Room, context echo.Context) *Participant {
	return &Participant{
		handler: NewProtocol().Switch(context),
		message: make(chan *packet.Message),
		room:    room,
		id:      id,
	}
}

func (participant *Participant) connect() {
	participant.room.register <- participant
}

func (participant *Participant) disconnect() {
	participant.handler.Close()
	participant.handler.SetCloseHandler(func(int, string) error {
		participant.room.unregister <- participant
		return nil
	})
}

func (participant *Participant) send() {
	var (
		duration time.Duration = time.Second * 60
		packet   *packet.Message
		err      error
	)
	participant.handler.SetReadLimit(4096)
	participant.handler.SetReadDeadline(time.Now().Add(duration))
	participant.handler.SetPongHandler(func(string) error {
		participant.handler.SetReadDeadline(time.Now().Add(duration))
		return nil
	})
	defer participant.disconnect()
	for {
		err = participant.handler.ReadJSON(&packet)
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
		participant.room.message <- packet
	}
}

func (participant *Participant) recv() {
	var (
		duration time.Duration = time.Second * 55
		ticker   *time.Ticker  = time.NewTicker(duration)
		err      error
	)
	defer participant.disconnect()
	defer ticker.Stop()
	for {
		duration := time.Second * 15
		select {
		case message := <-participant.message:
			participant.handler.SetWriteDeadline(time.Now().Add(duration))
			err = participant.handler.WriteJSON(message)
			if err != nil {
				logrus.Error("[participant.recv]: ", err.Error())
				return
			}
		case <-ticker.C:
			participant.handler.SetWriteDeadline(time.Now().Add(duration))
			err = participant.handler.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				return
			}
		}
	}
}
