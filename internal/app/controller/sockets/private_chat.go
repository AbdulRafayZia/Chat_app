package sockets

import (
	"fmt"

	"github.com/AbdulRafayZia/Gorilla-mux/utils"
)

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
