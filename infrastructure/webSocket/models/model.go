package models

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type RoomRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RoomResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Hub struct {
	Rooms      map[string]*Room `json:"rooms"`
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 10),
	}
}
