package chat

import (
	"net/http"

	"github.com/yesleymiranda/go-websocket/server/jwtoken"

	"github.com/yesleymiranda/go-toolkit/webapplication"
)

func Bind(app *webapplication.App, chatHub *Hub) {
	jwtSvc := jwtoken.NewService()
	svc := NewService(jwtSvc)
	routes(app, chatHub, svc)
}

func routes(app *webapplication.App, hub *Hub, svc Service) {
	app.Router.HandleFunc("/", GetChatHandler()).Methods(http.MethodGet)
	app.Router.HandleFunc("/ws/chat", func(w http.ResponseWriter, r *http.Request) { svc.Serve(hub, w, r) })
}
