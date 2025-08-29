package main

import (
	"database/sql"

	"github.com/bitkobie/todos/model"
)

type TodoRepo struct {
	db *sql.DB
}

func NewTodoRepo(db *sql.DB) model.TodoRepository {
	return &TodoRepo{
		db: db,
	}
}

// Create implements TodoStore.
func (t *TodoRepo) Create(todo *model.Todo) error {
	query := `INSERT INTO todos (title, complete) VALUES (NULLIF(?,''), ?)`

	result, err := t.db.Exec(query, todo.Title, todo.Complete)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	todo.Id = id

	return nil
}

// Update implements TodoStore.
func (t *TodoRepo) Update(todo *model.Todo) error {
	query := `UPDATE todos SET title = NULLIF(?,''), complete = ? WHERE id = ?`

	result, err := t.db.Exec(query, todo.Title, todo.Complete, todo.Id)
	if err != nil {
		return err
	}

	if rowsAffedted, _ := result.RowsAffected(); rowsAffedted == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// Get implements TodoStore.
func (t *TodoRepo) GetAll() ([]model.Todo, error) {
	query := `SELECT id, title, complete FROM todos`

	rows, err := t.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []model.Todo
	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&todo.Id, &todo.Title, &todo.Complete); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil

}

// Delete implements TodoStore.
func (t *TodoRepo) Delete(id int64) error {
	query := `DELETE FROM todos WHERE id = ?`

	result, err := t.db.Exec(query, id)
	if err != nil {
		return err
	}

	if rowsAffedted, _ := result.RowsAffected(); rowsAffedted == 0 {
		return sql.ErrNoRows
	}

	return nil
}
