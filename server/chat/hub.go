package chat

import (
	"github.com/yesleymiranda/go-websocket/server/account"
	"github.com/yesleymiranda/go-websocket/server/message"
	"github.com/yesleymiranda/go-websocket/server/room"
)

func New() *Hub {
	return &Hub{
		Rooms:           room.Initial(),
		register:        make(chan *account.Account),
		unregister:      make(chan *account.Account),
		broadcast:       make(chan *message.Message),
		broadcastSystem: make(chan *message.Message),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case acc := <-h.register:
			h.Rooms["A"].Accounts[acc.Token] = acc
		case acc := <-h.unregister:
			h.Rooms["A"].Accounts[acc.Username] = acc
		case msg := <-h.broadcast:
			for _, client := range h.Rooms["A"].Accounts {
				select {
				case client.Send <- msg.JSONStringfy():
				default:
					close(client.Send)
					delete(h.Rooms["A"].Accounts, client.Token)
				}
			}
		case m := <-h.broadcastSystem:
			m.Acc.Send <- m.JSONStringfy()
		}
	}
}
