package db

import (
	"database/sql"
	"os"

	_ "github.com/mxk/go-sqlite/sqlite3"
	"github.com/termith/minimblog/engine/common/logging"
)

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
