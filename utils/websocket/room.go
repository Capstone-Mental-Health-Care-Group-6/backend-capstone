package websocket

import (
	"FinalProject/utils/websocket/packet"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

type Room struct {
	clients map[*Client]bool
	join    chan *Client
	message chan *packet.Message
	sign    int
}

func NewRoom(sign int, clients ...*Client) *Room {
	room := &Room{
		clients: func() map[*Client]bool {
			buffer := make(map[*Client]bool)
			for _, client := range clients {
				buffer[client] = true
			}
			return buffer
		}(),
		join:    make(chan *Client),
		message: make(chan *packet.Message),
		sign:    sign,
	}
	for _, client := range clients {
		client.rooms[sign] = room
	}
	return room
}

func (r *Room) Join(client *Client) {
	r.clients[client] = true
	client.rooms[r.sign] = r
}

func (r *Room) Foward(message *packet.Message) {
	for client := range r.clients {
		result, err := base64.RawStdEncoding.DecodeString(client.sign)
		if err != nil {
			fmt.Println(err.Error())
		}
		sign, err := strconv.Atoi(strings.Split(string(result), "@")[1])
		if err != nil {
			fmt.Println(err.Error())
		}
		if message.To == sign {
			client.message <- message
		}
	}
}

func (r *Room) Listen() {
	for {
		select {
		case client := <-r.join:
			r.Join(client)
		case message := <-r.message:
			r.Foward(message)
		}
	}
}
