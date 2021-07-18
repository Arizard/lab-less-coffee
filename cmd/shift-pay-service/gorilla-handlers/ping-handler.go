package gorilla_handlers

import "net/http"

func PingHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("pong"))
}
