package main

import (
	"lemon/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HandlerIndex)
	http.ListenAndServe(":8089", nil)
}
