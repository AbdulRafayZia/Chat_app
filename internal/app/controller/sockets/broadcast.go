package sockets

import (
	"fmt"

	"github.com/AbdulRafayZia/Gorilla-mux/utils"
)

func (s *Server) BroadcastToAll(sender *utils.User, msg string) {
	message := utils.Message{
		Recipient: "broadcast",
		Content:   fmt.Sprintf("(Broadcast) %s: %s", sender.Username, msg),
	}

	// Loop through all connected users
	for _, u := range s.Client {
		// Skip sending the message to the sender
		if u.UserID == sender.UserID {
			continue
		}

		// Send the message to each connected user
		err := u.Connection.WriteJSON(message.Content)
		if err != nil {
			fmt.Println("Error sending broadcast message to", u.Username, ":", err)
		}
	}
}
