package routes

import (
	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/controller"
	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/controller/sockets"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func Routes() *mux.Router {
	r := mux.NewRouter()

	socketServer := sockets.Server{}
	socketServer.InitilzeServer()

	r.HandleFunc("/ws", socketServer.HandleConnections).Methods("GET")

	// go socketServer.HandleMessages()

	r.HandleFunc("/user/login", controller.LoginHandler).Methods("POST")

	r.HandleFunc("/user/signup", controller.CreateUserHandler).Methods("POST")
	r.HandleFunc("/staff/staffLogin", controller.StaffLogin).Methods("POST")
	r.HandleFunc("/create-room", controller.CreateRoomHandler).Methods("POST")

	// Join a room
	r.HandleFunc("/join-room", controller.JoinRoomHandler ).Methods("POST")

	// Room chat
	r.HandleFunc("/room/{roomID}",controller.RoomChatHandler ).Methods("POST")

	// Broadcast
	// r.HandleFunc("/broadcast", controller.BroadcastHandler).Methods("POST")

	// Private chat
	r.HandleFunc("/private/{recipientID}", controller.PrivateChatHandler).Methods("POST")
	r.HandleFunc("/refresh_token", controller.RefreshToken).Methods("POST")

	return r

}
