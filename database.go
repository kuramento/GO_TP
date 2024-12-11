package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDatabase() {
	var err error
	db, err = sql.Open("sqlite3", "tasks.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		status TEXT DEFAULT 'pending'
	);`

	if _, err := db.Exec(createTableQuery); err != nil {
		log.Fatalf("Failed to create table: %s", err)
	}
	log.Println("Database initialized successfully.")
}
