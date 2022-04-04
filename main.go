package main

import (
	"github.com/yesleymiranda/go-toolkit/webapplication"
	"github.com/yesleymiranda/go-websocket/server/chat"
	"github.com/yesleymiranda/go-websocket/server/login"
)

const port = "8081"

func main() {
	app := webapplication.New(&webapplication.ApplicationConfig{
		Port:     port,
		WithPing: true,
	})
	app.Initialize()

	hub := chat.New()
	go hub.Run()

	wireups(app, hub)

	_ = app.ListenAndServe()
}

func wireups(app *webapplication.App, hub *chat.Hub) {
	chat.Bind(app, hub)
	login.Bind(app, hub)
}
