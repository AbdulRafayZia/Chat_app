// var clients = make(map[*websocket.Conn]bool)
// var broadcast = make(chan MassageAlongSender)
package sockets

import (
	"github.com/AbdulRafayZia/Gorilla-mux/utils"
	"github.com/gorilla/websocket"
)

type BroadCast struct {
	Sender  *websocket.Conn
	Message utils.Message
}

type Server struct {
	Client      map[*websocket.Conn]string
	PrivateChat chan utils.Message
	Randoms     chan BroadCast
}

func (s *Server) InitilzeServer() {
	s.Client = make(map[*websocket.Conn]string)
	s.PrivateChat = make(chan utils.Message)
	s.Randoms = make(chan BroadCast)
}
