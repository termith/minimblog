package main

import (
	"fmt"
	"net/http"

	"github.com/termith/minimblog/common/config"
	"github.com/termith/minimblog/engine/handlers"
)

func main() {
	configuration := config.LoadConfig("/etc/minimblog/config.json")
	addr := fmt.Sprintf(":%d", configuration.Application.Port)

	http.HandleFunc("/post/", handlers.PostHandler)
	http.HandleFunc("/", handlers.RootHandler)
	http.ListenAndServe(addr, nil)
}
