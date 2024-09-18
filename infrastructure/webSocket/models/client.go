package models

import (
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	Socket   *websocket.Conn
	Message  chan *Message
	ID       string `json:"id"`
	RoomID   string `json:"roomId"`
	UserName string `json:"username"`
}

type ClientResponse struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
}

type Message struct {
	RoomID   string `json:"roomId"`
	UserName string `json:"username"`
	Content  string `json:"content"`
}

func (c *Client) Read(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		err := c.Socket.Close()
		if err != nil {
			return
		}
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		msg := &Message{
			RoomID:   c.RoomID,
			UserName: c.UserName,
			Content:  string(message),
		}

		hub.Broadcast <- msg
	}
}

func (c *Client) Write() {
	defer func() {
		err := c.Socket.Close()
		if err != nil {
			return
		}
	}()

	for {
		message, ok := <-c.Message
		if !ok {
			return
		}
		err := c.Socket.WriteJSON(message)
		if err != nil {
			return
		}
	}
}
