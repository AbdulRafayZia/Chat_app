package sockets

import "fmt"

func (s *Server) HandleMessages() {
	for {

		Randoms := <-s.Randoms
		for client := range s.Client {
			if client == Randoms.Sender {
				continue

			}
			err := client.WriteJSON(Randoms.Message)
			if err != nil {
				fmt.Println(err)
				client.Close()
				delete(s.Client, client)
			}

		}
		PrivateChat := <-s.PrivateChat
		for client, username := range s.Client {
			if username != "" {
				if username == PrivateChat.Receiver {
					continue
				}
				err := client.WriteJSON(PrivateChat.Message)
				if err != nil {
					fmt.Println(err)
					client.Close()
					delete(s.Client, client)
				}
			} else {

				err := client.WriteJSON(PrivateChat.Message)
				if err != nil {
					fmt.Println(err)
					client.Close()
					delete(s.Client, client)
				}
			}

		}
	}
}
