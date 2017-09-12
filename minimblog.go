package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/termith/minimblog/common/config"
	"github.com/termith/minimblog/common/db"
	"github.com/termith/minimblog/common/logging"
	"github.com/termith/minimblog/engine/handlers"
)

var mode string

func startServer(port int) {
	addr := fmt.Sprintf(":%d", port)
	http.HandleFunc("/post/", handlers.PostHandler)
	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/ping", handlers.PingHandler)

	http.ListenAndServe(addr, nil)
}

func main() {

	flag.StringVar(&mode, "mode", "development", "Run mode: production or development")
	flag.Parse()

	configuration := config.LoadConfig("resources/config.json")

	db.InitDBConnection(configuration.Database.Driver, configuration.Database.Url)

	logging.InitLogging("minimblog.log", mode)

	startServer(configuration.Application.Port)
}
