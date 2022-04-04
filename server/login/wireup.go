package login

import (
	"net/http"

	"github.com/yesleymiranda/go-websocket/server/jwtoken"

	"github.com/yesleymiranda/go-websocket/server/chat"

	"github.com/yesleymiranda/go-toolkit/webapplication"
)

func Bind(app *webapplication.App, chatHub *chat.Hub) {
	jwtSvc := jwtoken.NewService()
	svc := NewService(chatHub, jwtSvc)
	routes(app, svc)
}

func routes(app *webapplication.App, svc Service) {
	app.Router.HandleFunc("/login", MakeGetLoginHandler()).Methods(http.MethodGet)
	app.Router.HandleFunc("/login", MakePostLoginHandler(svc)).Methods(http.MethodPost)
}
