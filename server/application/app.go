package application

import (
	"net/http"
)

type App struct {
	Port string
}

func New(port string) *App {
	return &App{
		Port: port,
	}
}

func Index() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fileName := "server/application/index.html"
		http.ServeFile(w, r, fileName)
	})
}
