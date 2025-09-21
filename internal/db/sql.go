package db

import (
	"database/sql"
	"log"

	_ "github.com/glebarez/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite", "./forum.db")
	if err != nil {
		log.Fatal("Error open database:", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("Error connect to database:", err)
	}

}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
