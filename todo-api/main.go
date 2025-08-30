package main

import (
	"database/sql"
	"log"
	"net/http"

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

	repo := NewTodoRepo(db)
	handler := NewHandler(repo)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /todos", handler.CreateTodo)
	mux.HandleFunc("PUT /todos", handler.UpdateTodo)
	mux.HandleFunc("GET /todos", handler.GetAllTodos)
	mux.HandleFunc("DELETE /todos/{id}", handler.DeleteTodo)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Printf("server running on port %s", server.Addr)
	log.Fatal(server.ListenAndServe())
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
