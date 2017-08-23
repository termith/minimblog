package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/termith/minimblog/common/config"
	"github.com/termith/minimblog/common/logging"
	"github.com/termith/minimblog/engine/handlers"
)

var mode string

func main() {

	flag.StringVar(&mode, "mode", "development", "Run mode: production or development")
	flag.Parse()

	if mode != "production" || mode != "development" {
		fmt.Println("Mode should be one of (production, development)")
		os.Exit(1)
	}
	configuration := config.LoadConfig("/etc/minimblog/config.json")

	logging.InitLogs(configuration, mode)
	addr := fmt.Sprintf(":%d", configuration.Application.Port)
	http.HandleFunc("/post/", handlers.PostHandler)
	http.HandleFunc("/", handlers.RootHandler)

	http.ListenAndServe(addr, nil)
}
