package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "file:todos.db")
	if err != nil {
		log.Fatal(err)
	}
	if err := initDB(db); err != nil {
		log.Fatal(err)
	}
	_ = NewTodoRepo(db)

	fmt.Println("Hello World")
}

func initDB(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		complete BOOLEAN NOT NULL DEFAULT FALSE
	);`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
