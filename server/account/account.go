package account

import (
	"github.com/gorilla/websocket"
)

type Account struct {
	Token    string `json:"token,omitempty"`
	Username string `json:"username,omitempty"`
	Conn     *websocket.Conn
	Send     chan []byte
}
