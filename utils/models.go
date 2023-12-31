package utils

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Message struct {
	RoomName  string `json:"room_name"`
	Content   string `json:"Content"`
	Recipient string `json:"recipient"`
	Type      string `json:"type"`
}
type Token struct {
	AccessToken string `json:"acessToken"`
	RefreshToken string `json:"refreshToken"`

}
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ID       int    `json:"id"`
	Role     string `json:"role"`

}


type Room struct {
	RoomID   string
	RoomName string
	Users    map[string]*User
	Admins   map[string]bool
}
type User struct {
	UserID     uint
	Username   string
	Connection *websocket.Conn
	Rooms      map[string]struct{} // User's joined rooms
	Mu         sync.Mutex
}

type Roooms struct {
	Room Room
}
type RoomMessages struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}
