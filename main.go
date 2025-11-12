package main

import (
	"EmqxBackEnd/database"
	"EmqxBackEnd/router"
	"database/sql"
)

func main() {
	database.Init()
	defer func(DB *sql.DB) {
		_ = DB.Close()
	}(database.DB)
	r := router.Setup()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
