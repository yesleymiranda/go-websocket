package chat

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/yesleymiranda/go-toolkit/logger"
	"github.com/yesleymiranda/go-websocket/server/account"
	"github.com/yesleymiranda/go-websocket/server/jwtoken"
	"github.com/yesleymiranda/go-websocket/server/message"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 5 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline  = []byte{'\n'}
	space    = []byte{' '}
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Service interface {
	Serve(chatHub *Hub, w http.ResponseWriter, r *http.Request)
}

type service struct {
	jwtSvc jwtoken.Service
}

func NewService(jwtSvc jwtoken.Service) Service {
	return &service{
		jwtSvc,
	}
}

func (s *service) Serve(chatHub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("error on chat serve", err)
		return
	}

	t := r.URL.Query().Get("t")
	acc, err := s.jwtSvc.ReadAndValidate(t)
	if err != nil {
		http.Error(w, "token is required", http.StatusUnauthorized)
		return
	}

	client := accountRegister(chatHub, acc, conn)
	go readMessage(chatHub, client)
	go writeMessage(chatHub, client)
}

func accountRegister(chatHub *Hub, acc *account.Account, conn *websocket.Conn) *account.Account {
	acc.Send = make(chan []byte, 256)
	acc.Conn = conn
	chatHub.register <- acc

	msg := fmt.Sprintf("[action:register][room:%s][username:%s]", "A", acc.Username)
	logger.Info(msg)

	return acc
}

func readMessage(chatHub *Hub, acc *account.Account) {
	defer func() {
		chatHub.unregister <- acc
		msg := fmt.Sprintf("[action:unregister][room:%s][username:%s]", "A", acc.Username)
		logger.Info(msg)
		acc.Conn.Close()
	}()
	acc.Conn.SetReadLimit(maxMessageSize)
	// c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	// c.Conn.SetPongHandler(func(string) error { c.Conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, messageReceived, err := acc.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		msg := string(bytes.TrimSpace(bytes.Replace(messageReceived, newline, space, -1)))
		m := message.New(acc, msg)
		if m.IsError() {
			chatHub.broadcastSystem <- m
		} else {
			chatHub.broadcast <- m
		}

		msgLog := fmt.Sprintf("[action:broadcast][room:%s][user:%s][message:%s]", "A", acc.Username, m.Message)
		logger.Info(msgLog)

	}
}

func writeMessage(chatHub *Hub, c *account.Account) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
