package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func Database() *sql.DB {
	loadEnv()

	dbPath := os.Getenv("DBPATH")

	var err error

	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
