package login

import "github.com/yesleymiranda/go-websocket/server/account"

type Service interface {
	Login(account *account.Account) (string, error)
}
