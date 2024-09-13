package models

import "golang.org/x/net/websocket"

type Client struct {
	ID     string `json:"id"`
	Socket *websocket.Conn
}

type Message struct {
	UserName string `json:"user_name"`
	Content  string `json:"content"`
}

type ClientMessage struct {
}
