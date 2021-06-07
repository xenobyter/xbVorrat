package api

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	
	"log"
	"os"
)

func init() {
	var err error
	home, _ := os.UserHomeDir()
	db, err = sql.Open("sqlite3", home+"/.xbVorrat")
	if err != nil {
		log.Panic(err)
	}
	createTables()
}