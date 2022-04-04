package message

import (
	"encoding/json"
	"github.com/rs/zerolog"
	"github.com/yesleymiranda/go-websocket/server/account"
)

const messageError = "menssage is required"
const system = "system"

type Message struct {
	User    string           `json:"u,omitempty"`
	Message string           `json:"m,omitempty"`
	Level   string           `json:"l,omitempty"`
	Acc     *account.Account `json:"-"`
}

func New(acc *account.Account, msg string) *Message {
	if msg == "" {
		m := NewIsRequiredError(acc)
		return m
	}
	return NewInfo(acc.Username, msg)
}

func NewIsRequiredError(acc *account.Account) *Message {
	return &Message{User: system, Level: zerolog.LevelErrorValue, Message: messageError, Acc: acc}
}

func NewInfo(user string, msg string) *Message {
	return &Message{User: user, Level: zerolog.LevelInfoValue, Message: msg}
}

func (m *Message) IsError() bool {
	return m.Level == zerolog.LevelErrorValue
}

func (m *Message) JSONStringfy() []byte {
	j, _ := json.Marshal(m)
	return j
}
