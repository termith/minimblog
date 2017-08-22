package main

import (
	"net/http"

	"github.com/termith/minimblog/engine/handlers"
)

func main() {
	http.HandleFunc("/post/", handlers.PostHandler)
	http.HandleFunc("/", handlers.RootHandler)
	http.ListenAndServe(":8080", nil)
}
