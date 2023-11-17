package websocket

import (
	"fmt"

	"FinalProject/utils/websocket/packet"
)

type Room struct {
	participants map[*Participant]bool
	message      chan *packet.Message
	register     chan *Participant
	unregister   chan *Participant
	id           int
}

func NewRoom(id int) *Room {
	return &Room{
		participants: make(map[*Participant]bool),
		message:      make(chan *packet.Message),
		register:     make(chan *Participant),
		unregister:   make(chan *Participant),
		id:           id,
	}
}

func (room *Room) join(participant *Participant) {
	if _, found := room.participants[participant]; !found {
		fmt.Println(participant.id, "join")
		room.participants[participant] = true
	}
}

func (room *Room) leave(participant *Participant) {
	if _, found := room.participants[participant]; found {
		close(participant.message)
		delete(room.participants, participant)
	}
}

func (room *Room) foward(message *packet.Message) {
	for participant := range room.participants {
		if message.To == participant.id {
			participant.message <- message
		}
	}
}

func (room *Room) Listen() {
	go func() {
		for {
			select {
			case participant := <-room.register:
				room.join(participant)
			case participant := <-room.unregister:
				room.leave(participant)
			case message := <-room.message:
				room.foward(message)
			}
		}
	}()
}

func (room *Room) Serve(participant *Participant) {
	participant.connect()
	go participant.recv()
	go participant.send()
}
