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

	go socketServer.HandleMessages()

	r.HandleFunc("/user/login", controller.LoginHandler).Methods("POST")
	r.HandleFunc("/fileProcess", controller.ProcessFile).Methods("POST")
	r.HandleFunc("/user/signup", controller.CreateUserHandler).Methods("POST")
	r.HandleFunc("/staff/staffLogin", controller.StaffLogin).Methods("POST")
	r.HandleFunc("/user/user_processes", controller.GetUsersProcessses).Methods("GET")
	r.HandleFunc("/user/get_process/{id}", controller.GetProcessById).Methods("GET")
	r.HandleFunc("/staff/statistics", controller.Statistics).Methods("POST")
	r.HandleFunc("/staff/get_all_processes", controller.GetAllProcesses).Methods("GET")
	r.HandleFunc("/staff/get_processes_by_username", controller.GetProcessByName).Methods("POST")
	r.HandleFunc("/refresh_token", controller.RefreshToken).Methods("POST")

	return r

}
