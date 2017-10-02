package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/termith/minimblog/common/db"
	"github.com/termith/minimblog/engine/common/config"
	"github.com/termith/minimblog/engine/common/logging"
	"github.com/termith/minimblog/engine/handlers"
)

var mode string
var configPath string

func startServer(port int) {
	addr := fmt.Sprintf(":%d", port)
	http.HandleFunc("/post/", handlers.PostHandler)
	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/ping", handlers.PingHandler)

	http.ListenAndServe(addr, nil)
}

func main() {

	flag.StringVar(&mode, "mode", "development", "Run mode: production or development")
	flag.StringVar(&configPath, "config", "resources/config.json", "Path to config")
	flag.Parse()

	configuration := config.LoadConfig(configPath)

	db.InitDBConnection(configuration.Database.Driver, configuration.Database.Url)

	logging.InitLogging("minimblog.log", mode)

	startServer(configuration.Application.Port)
}
