package websocket

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Server struct {
	clients     map[string]*Client
	client_refs map[string]string
	rooms       map[int]*Room
}

func NewServer() *Server {
	return &Server{
		clients:     make(map[string]*Client),
		client_refs: make(map[string]string),
		rooms:       make(map[int]*Room),
	}
}

func (s *Server) FindClient(ref string) *Client {
	sign := s.client_refs[ref]
	return s.clients[sign]
}

func (s *Server) FindRoom(sign int) *Room {
	return s.rooms[sign]
}

func (s *Server) CreateClient(ctx echo.Context, user int, role string) string {
	ref, client := NewClient(ctx, s, user, role)
	s.clients[client.sign] = client
	s.client_refs[ref] = client.sign
	go s.clients[client.sign].Send()
	go s.clients[client.sign].Recv()
	return client.sign
}

func (s *Server) CreateRoom(sign int, refs ...string) *Room {
	clients := func() (buffer []*Client) {
		for _, ref := range refs {
			buffer = append(buffer, s.FindClient(ref))
		}
		return buffer
	}()
	s.rooms[sign] = NewRoom(sign, clients...)
	go s.rooms[sign].Listen()
	return s.rooms[sign]
}

func (s *Server) JoinRoom(sign int, ref string) {
	s.rooms[sign].join <- s.FindClient(ref)
}

func (s *Server) DeleteClient(sign string) {
	for i, room := range s.clients[sign].rooms {
		logrus.Infof("[ws.server]: client@%s keluar dari room %d", sign, i)
		room.leave <- s.clients[sign]
	}
	delete(s.clients, sign)
}

func (s *Server) DeleteRoom(sign int) {
	for client := range s.rooms[sign].clients {
		logrus.Infof("[ws.server]: client@%s keluar dari room %d", client.sign, sign)
		s.rooms[sign].leave <- client
	}
	// close(s.rooms[sign].join)
	// close(s.rooms[sign].leave)
	close(s.rooms[sign].message)
	delete(s.rooms, sign)
}
