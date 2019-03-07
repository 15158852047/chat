package models

import "github.com/gorilla/websocket"

type HelloMsg struct {
	Name string `json:"name"`
	Time string `json:"time"`
	Ws *websocket.Conn
}
