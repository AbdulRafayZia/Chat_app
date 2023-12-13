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

	r.HandleFunc("/refresh-token", controller.RefreshToken).Methods("POST")

	return r

}
