package chat

import (
	"github.com/yesleymiranda/go-websocket/server/account"
	"github.com/yesleymiranda/go-websocket/server/message"
	"github.com/yesleymiranda/go-websocket/server/room"
)

type Hub struct {
	Rooms           map[string]*room.Room
	register        chan *account.Account
	unregister      chan *account.Account
	broadcast       chan *message.Message
	broadcastSystem chan *message.Message
}
