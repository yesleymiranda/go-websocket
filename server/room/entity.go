package room

import "github.com/yesleymiranda/go-websocket/server/account"

type Room struct {
	Accounts map[string]*account.Account
}
