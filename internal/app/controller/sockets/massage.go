package sockets

// import "fmt"

// func (s *Server) HandleMessages() {
// 	for {

// 		PrivateChat := <-s.PrivateChat
// 		for client, username := range s.Client {
// 			if PrivateChat.Receiver != "" {
// 				if username == PrivateChat.Receiver {
// 					err := client.WriteJSON(PrivateChat.Message)
// 					if err != nil {
// 						fmt.Println(err)
// 						client.Close()
// 						delete(s.Client, client)
// 					}
// 				}

// 			} else {

// 				err := client.WriteJSON(PrivateChat.Message)
// 				if err != nil {
// 					fmt.Println(err)
// 					client.Close()
// 					delete(s.Client, client)
// 				}
// 			}

// 		}
// 	}
// }
