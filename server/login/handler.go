package login

import (
	"html/template"
	"net/http"

	"github.com/yesleymiranda/go-websocket/server/account"
	"github.com/yesleymiranda/go-websocket/toolkit/request"

	"github.com/yesleymiranda/go-toolkit/web"
)

func MakeGetLoginHandler() web.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("public/login/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		t.Execute(w, nil)
	}
}

func MakePostLoginHandler(svc Service) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		body := &account.Account{}
		err := request.Unmarshal(r, body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res, err := svc.Login(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		web.EncodeJSON(w, res, http.StatusOK)
	}
}
