package utils

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Receiver string `json:"receiver"`
}

type Room struct {
	Name   string   `json:"name"`
	Users  []string `json:"users"`
	Admins []string `json:"admins"`
	Messages RoomMessages 
}
type Roooms struct{
	Room Room
}
type RoomMessages struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}
