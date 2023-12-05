package websocket

import (
	"github.com/labstack/echo/v4"
)

type Server struct {
	clients map[int]*Client
	rooms   map[int]*Room
}

func NewServer() *Server {
	return &Server{
		clients: make(map[int]*Client),
		rooms:   make(map[int]*Room),
	}
}

func (s *Server) FindClient(sign int) *Client {
	return s.clients[sign]
}

func (s *Server) FindRoom(sign int) *Room {
	return s.rooms[sign]
}

func (s *Server) CreateClient(ctx echo.Context, sign int) *Client {
	s.clients[sign] = NewClient(ctx, s, sign)
	go s.clients[sign].Send()
	go s.clients[sign].Recv()
	return s.clients[sign]
}

func (s *Server) CreateRoom(sign int, users ...int) *Room {
	clients := func() (buffer []*Client) {
		for _, user := range users {
			buffer = append(buffer, s.FindClient(user))
		}
		return buffer
	}()
	s.rooms[sign] = NewRoom(sign, clients...)
	go s.rooms[sign].Listen()
	return s.rooms[sign]
}

func (s *Server) JoinRoom(sign int, a int) {
	s.rooms[sign].join <- s.FindClient(a)
}

func (s *Server) DeleteClient(sign int) {
	delete(s.clients, sign)
}

func (s *Server) DeleteRoom(sign int) {
	// close(s.rooms[sign].join)
	// close(s.rooms[sign].message)
	delete(s.rooms, sign)
}
