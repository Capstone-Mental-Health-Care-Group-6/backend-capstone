package websocket

type Server struct {
	rooms map[int]*Room
}

func NewServer() *Server {
	return &Server{
		rooms: make(map[int]*Room),
	}
}

func (server *Server) CreateRoom(id int) {
	if server.rooms[id] == nil {
		server.rooms[id] = NewRoom(id)
	}
}

func (server *Server) FindRoom(id int) *Room {
	return server.rooms[id]
}

func (server *Server) DeleteRoom(id int) {
	if server.rooms[id] != nil {
		delete(server.rooms, id)
	}
}
