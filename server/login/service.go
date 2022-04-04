package login

import (
	"errors"
	"fmt"

	"github.com/yesleymiranda/go-websocket/server/jwtoken"

	"github.com/yesleymiranda/go-toolkit/logger"

	"github.com/yesleymiranda/go-websocket/server/account"
	"github.com/yesleymiranda/go-websocket/server/chat"
)

type service struct {
	chatHub    *chat.Hub
	jwtService jwtoken.Service
}

func NewService(chatHub *chat.Hub, jwtService jwtoken.Service) Service {
	return &service{
		chatHub,
		jwtService,
	}
}

func (s service) Login(acc *account.Account) (string, error) {
	if acc.Username == "" {
		return "", errors.New("error on login")
	}

	token, err := s.jwtService.Create(acc)
	if err != nil {
		return "", errors.New("error trying to generate token:" + err.Error())
	}

	_, err = s.jwtService.ReadAndValidate(token)
	if err != nil {
		return "", errors.New("error when trying to validate the token:" + err.Error())
	}

	s.chatHub.Rooms["A"].Accounts[token] = acc

	msg := fmt.Sprintf("[action:login][username:%s]", acc.Username)
	logger.Info(msg)

	return token, nil
}
