package sockets

import (
	"fmt"

	"github.com/AbdulRafayZia/Gorilla-mux/utils"
)

var rooms = make(map[string]map[*utils.User]struct{})

func(s *Server) BroadcastToRoom(sender *utils.User, msg utils.Message) {
	// Send the message to all users in the room
	usersInRoom, exists := s.Rooms[msg.RoomName]
	if !exists {
		return
	}

	message := utils.Message{
		Recipient:    "RoomChat",
		Content: fmt.Sprintf("%s: %s", sender.Username, msg.Content),
		RoomName:  msg.RoomName,
	}

	for user := range usersInRoom.Users {
		if user.Username == sender.Username {
			continue // Skip sending the message to the sender
		}
		err := user.Connection.WriteJSON(message.Content)
		if err != nil {
			fmt.Println("Error sending broadcast message to", user.Username, ":", err)
		}
	}
}
func (s *Server) JoinRoom(user *utils.User, roomName string) {
    user.Mu.Lock()
    defer user.Mu.Unlock()

    // Leave existing rooms before joining the new room
    // LeaveAllRooms(user)

    // Join the room
  
    room,roomExists := s.Rooms[roomName]
    if !roomExists {
        if user.Connection != nil {
            user.Connection.WriteJSON("room cannot exist")
        } else {
            fmt.Println("User connection is nil. Cannot send room cannot exist message.")
        }
        return
    }

    room.Users[user] = utils.Message{
        Recipient: "notification",
        Content:   fmt.Sprintf("You have joined room %s", roomName),
        RoomName:  roomName,
    }

    // Notify the user that they have joined the room
    if user.Connection != nil {
        err := user.Connection.WriteJSON(room.Users[user].Content)
        if err != nil {
            fmt.Println("Error sending room join notification to", user.Username, ":", err)
        }
    } else {
        fmt.Println("User connection is nil. Cannot send room join notification.")
    }
}
func (s *Server) CreateRoom(user *utils.User, roomName string) {
    user.Mu.Lock()
    defer user.Mu.Unlock()

    // Check if the room already exists in s.Rooms
    room, roomExists := s.Rooms[roomName]
    if !roomExists {
        // If the room doesn't exist, create it
        room = &Members{
            Users: make(map[*utils.User]utils.Message),
        }
        s.Rooms[roomName] = room
    }

    // Add the user to the room
    room.Users[user] = utils.Message{
        Recipient: "notification",
        Content:   fmt.Sprintf("You have created and joined room %s", roomName),
        RoomName:  roomName,
    }

    // Send room creation notification to the user
    err := user.Connection.WriteJSON(room.Users[user].Content)
    if err != nil {
        fmt.Println("Error sending room creation notification to", user.Username, ":", err)
    }
}
	

func LeaveAllRooms(user *utils.User) {
	user.Mu.Lock()
	defer user.Mu.Unlock()

	for roomName := range user.Rooms {
		delete(rooms[roomName], user)
		// Notify the user that they have left the room
		notification := utils.Message{
			Recipient:    "notification",
			Content: fmt.Sprintf("You have left room %s", roomName),
			RoomName:  roomName,
		}
		err := user.Connection.WriteJSON(notification)
		if err != nil {
			fmt.Println("Error sending room leave notification to", user.Username, ":", err)
		}
		delete(user.Rooms, roomName)
	}
}
