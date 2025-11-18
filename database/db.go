package database

import (
	"EmqxBackEnd/config"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("postgres", config.DBConnStr)
	if err != nil {
		log.Fatal("Failed to connect to DB", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("Failed to ping DB", err)
	}
}
