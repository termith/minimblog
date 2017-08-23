package logging

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/termith/minimblog/common/config"
)

var (
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
)

func InitLogs(cfg config.Config, mode string) {
	logFlags := log.Ldate | log.Ltime | log.Lshortfile
	logFile, err := os.OpenFile(cfg.Logging.File, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("Error while open logfile")
		os.Exit(1)
	}
	Info = log.New(os.Stdout, "INFO: ", logFlags)
	Warn = log.New(os.Stdout, "WARNING: ", logFlags)
	Error = log.New(logFile, "ERROR: ", logFlags)

	switch mode {
	case "production":
		Info.SetOutput(ioutil.Discard)
		Warn.SetOutput(ioutil.Discard)
	case "development":
		Info.SetOutput(logFile)
	}

}
