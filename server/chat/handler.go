package chat

import (
	"html/template"
	"net/http"

	"github.com/yesleymiranda/go-toolkit/web"
)

func GetChatHandler() web.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("public/chat/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		t.Execute(w, nil)
	}
}
