package handlers

import "net/http"

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Lemon!"))
}
