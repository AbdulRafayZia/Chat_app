// var clients = make(map[*websocket.Conn]bool)
// var broadcast = make(chan MassageAlongSender)
package sockets

import (
	"github.com/AbdulRafayZia/Gorilla-mux/utils"
)

type Server struct {
	Client map[string]*utils.User
	Rooms  map[string]*Members
}
type Members struct {
	
	Users    map[*utils.User]utils.Message
}

func (s *Server) InitilzeServer() {
	s.Client = make(map[string]*utils.User)
	s.Rooms=make(map[string]*Members)

	// s.Randoms = make(chan BroadCast)
}
