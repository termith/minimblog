package db

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/termith/minimblog/engine/common/logging"
)

// TODO: https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/05.3.html

var connection *sql.DB

func Connection() *sql.DB {
	return connection
}

func InitDBConnection(driver, url string) {
	var err error
	connection, err = sql.Open(driver, url)
	if err != nil {
		logging.Error("Error while init DB connection")
		os.Exit(1)
	}
}
