package main

import (
	"fmt"
	"net/http"

	"github.com/yesleymiranda/go-websocket/server/application"
	"github.com/yesleymiranda/go-websocket/server/login"
)

func main() {
	app := application.New("8081")

	application.Index()
	login.Login()

	err := http.ListenAndServe(fmt.Sprintf(":%s", app.Port), nil)
	if err != nil {
		panic("error o listen and serve")
	}
}
