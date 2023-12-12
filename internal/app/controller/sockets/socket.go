package sockets

import (
	"fmt"
	"net/http"

	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/validation"

	"github.com/AbdulRafayZia/Gorilla-mux/pkg/jwt"
	"github.com/AbdulRafayZia/Gorilla-mux/utils"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (s *Server) HandleConnections(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	// Assuming you have a function to get user details from authentication
	user, err := GetCurrentUser(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := utils.User{
		Username:   user.Username,
		UserID:     user.UserID,
		Connection: conn,
	}

	// Register the user
	s.Client[user.Username] = &utils.User{
		UserID:     user.UserID,
		Username:   user.Username,
		Connection: conn,
	}

	for {
		var msg utils.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			delete(s.Client, user.Username)
			return
		}

		switch msg.Type {
		case "broadcast":
			s.BroadcastToAll(&user, msg.Content)
		case "private":
			s.SendPrivateMessage(&user, msg)
		case "join_room":
			s.JoinRoom(&client, msg.RoomName)
		case "room_chat":
			s.BroadcastToRoom(&client, msg)
		case "create_room":
			s.CreateRoom(&client, msg.RoomName)
		}
	}
}

func GetCurrentUser(w http.ResponseWriter, r *http.Request) (utils.User, error) {
	tokenString, err := jwt.GetToken(w, r)
	if tokenString == "" || err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "could not provide autherization bearer", http.StatusUnauthorized)
		return utils.User{}, err

	}
	claims, err := validation.VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Could not Get Claims")
		return utils.User{}, err
	}
	user := utils.User{
		UserID:   claims.Id,
		Username: claims.Username,
	}
	return user, nil

}

func (s *Server) SendPrivateMessage(sender *utils.User, msg utils.Message) {

	message := utils.Message{
		Recipient: msg.Recipient, // Indicate that it's a private message from the sender
		Content:   fmt.Sprintf("(Private) %s: %s", sender.Username, msg.Content),
	}

	// Find the recipient user
	recipient, exists := s.Client[msg.Recipient]
	if !exists {
		fmt.Println("Recipient not found:", msg.Recipient)
		return
	}

	// Send the private message to the recipient
	err := recipient.Connection.WriteJSON(message.Content)
	if err != nil {
		fmt.Println("Error sending private message to", recipient.Username, ":", err)
	}
}

func (s *Server) BroadcastToAll(sender *utils.User, msg string) {
	message := utils.Message{
		Recipient: "broadcast",
		Content:   fmt.Sprintf("(Broadcast) %s: %s", sender.Username, msg),
	}

	// Loop through all connected users
	for _, u := range s.Client {
		// Skip sending the message to the sender
		if u.Username == sender.Username {
			continue
		}

		// Send the message to each connected user
		err := u.Connection.WriteJSON(message.Content)
		if err != nil {
			fmt.Println("Error sending broadcast message to", u.Username, ":", err)
		}
	}
}
