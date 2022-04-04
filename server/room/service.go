package room

import "github.com/yesleymiranda/go-websocket/server/account"

func Initial() map[string]*Room {
	rooms := make(map[string]*Room)
	rooms["A"] = &Room{make(map[string]*account.Account)}
	rooms["B"] = &Room{make(map[string]*account.Account)}
	return rooms
}
