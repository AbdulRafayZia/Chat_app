package controller

// func BroadcastHandler(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")

// 	var request utils.StatsRequest
// 	err := json.NewDecoder(r.Body).Decode(&request)
// 	if err != nil {
// 		http.Error(w, "unable to get data", http.StatusBadRequest)
// 		return
// 	}
// 	tokenString, err := jwt.GetToken(w, r)
// 	if tokenString == "" || err != nil {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		http.Error(w, "could not provide autherization bearer", http.StatusUnauthorized)
// 		return
// 	}

// 	claims, err := validation.VerifyToken(tokenString)
// 	if err != nil {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		http.Error(w, " could not get Claims ", http.StatusUnauthorized)

// 		return
// 	}
// 	// user := getCurrentUser(r)
// 	user := claims.Username

// 	// Read the message from the client
// 	var message struct {
// 		Message string `json:"message"`
// 	}

// 	err = json.NewDecoder(r.Body).Decode(&message)
// 	if err != nil {
// 		http.Error(w, "Invalid JSON", http.StatusBadRequest)
// 		return
// 	}

// 	// Broadcast the message to all connected clients
// 	broadcastToAll(user, message.Message)

// 	w.WriteHeader(http.StatusOK)
// }

// Assume you have a Message struct defined somewhere in your code
type Message struct {
	Recipient string `json:"recipient"`
	Content   string `json:"content"`
}

// func BroadcastToAll(sender *utils.User, content string) {
// 	message := Message{
// 		Recipient: "broadcast",
// 		Content:   fmt.Sprintf("(Broadcast) %s: %s", sender.Username, content),
// 	}

// 	// Loop through all connected users
// 	for _, u := range s.Client {
// 		// Skip sending the message to the sender
// 		if u.UserID == sender.UserID {
// 			continue
// 		}

// 		// Send the message to each connected user
// 		err := u.Connection.WriteJSON(message)
// 		if err != nil {
// 			fmt.Println("Error sending broadcast message to", u.Username, ":", err)
// 		}
// 	}
// }
