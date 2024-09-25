package storage

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() error {
	var err error
	db, err = sql.Open("sqlite3", "./task.db")

	if err != nil {
		log.Fatalf("Error opening the database:  %v", err)
		return err
	}

	err = createTable()
	if err != nil {
		log.Fatalf("Error creating the table: %v", err)
		return err
	}

	return nil
}

func createTable() error {
	createTaskTable := `
    CREATE TABLE IF NOT EXISTS tasks (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      name TEXT,
      status TEXT
    )
  `
	_, err := db.Exec(createTaskTable)
	return err
}

func GetDB() *sql.DB {
  if db == nil {
    log.Println("Error getting the database, connection is nil")
  }
	return db
}
