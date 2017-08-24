package logging

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	info  *log.Logger
	warn  *log.Logger
	err   *log.Logger
	debug *log.Logger
)

func Info(args ...interface{}) {
	info.Println(args...)
}

func Error(args ...interface{}) {
	err.Println(args...)
}

func Warn(args ...interface{}) {
	warn.Println(args...)
}

func Debug(args ...interface{}) {
	debug.Println(args...)
}

// InitLogging initializes logging subsystems
func InitLogging(logFilePath string, mode string) {
	logFlags := log.Ldate | log.Ltime | log.Lshortfile
	logFile, er := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE, 0755)
	if er != nil {
		fmt.Println("Error while open logfile: " + logFilePath)
		os.Exit(1)
	}
	info = log.New(os.Stdout, "INFO: ", logFlags)
	warn = log.New(os.Stdout, "WARNING: ", logFlags)
	err = log.New(logFile, "ERROR: ", logFlags)
	debug = log.New(os.Stdout, "DEBUG", logFlags)

	switch mode {
	case "production":
		info.SetOutput(ioutil.Discard)
		warn.SetOutput(ioutil.Discard)
		debug.SetOutput(ioutil.Discard)
	case "development":
		info.SetOutput(logFile)
	case "debug":
		err.SetOutput(os.Stdout)
	}

}
