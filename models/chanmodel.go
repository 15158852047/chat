package models

import (
	"github.com/gorilla/websocket"
	"time"
)

type HelloMsg struct {
	Name     string `json:"name"`
	Time     string `json:"time"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Ws       *websocket.Conn
}

type Chat struct {
	From string    `json:"from"`
	To   string    `json:"to"`
	Time time.Time `json:"time"`
	Msg  string    `json:"msg"`
	Wss  []*websocket.Conn
}

type Recive struct {
	Code string `json:"code"`
	To   string `json:"to"`
	Msg  string `json:"msg"`
}
