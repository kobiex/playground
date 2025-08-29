package main

import (
	"database/sql"
	"slices"
	"testing"

	"github.com/bitkobie/todos/model"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", "file:test.db")
	if err != nil {
		t.Fatalf("failed to open sqlite: %v", err)
	}
	err = initDB(db)
	if err != nil {
		t.Fatalf("failed to initialize db: %v", err)
	}

	return db
}

func TestCreate(t *testing.T) {
	testCases := []struct {
		name    string
		todo    *model.Todo
		wantErr bool
	}{
		{
			name:    "valid todo",
			todo:    &model.Todo{Title: "Run this test", Complete: false},
			wantErr: false,
		},
		{
			name:    "invalid todo: missing required field",
			todo:    &model.Todo{Complete: false, Title: ""},
			wantErr: true,
		},
		{
			name:    "valid todo: missing optional field",
			todo:    &model.Todo{Title: "leave status out"},
			wantErr: false,
		},
	}

	repo := NewTodoRepo(setupTestDB(t))

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.Create(tc.todo)
			if tc.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	repo := NewTodoRepo(setupTestDB(t))
	todo := &model.Todo{Title: "Initial Title", Complete: false}
	err := repo.Create(todo)
	if err != nil {
		t.Fatalf("failed to create initial todo: %v", err)
	}

	testCases := []struct {
		name    string
		todo    *model.Todo
		wantErr bool
	}{
		{
			name:    "valid todo",
			todo:    &model.Todo{Id: todo.Id, Title: "Run this test", Complete: true},
			wantErr: false,
		},
		{
			name:    "invalid todo: missing required field",
			todo:    &model.Todo{Id: todo.Id, Complete: false, Title: ""},
			wantErr: true,
		},
		{
			name:    "invalid todo: invalid id",
			todo:    &model.Todo{Id: -98, Title: "leave status out"},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.Update(tc.todo)
			if tc.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestGetAll(t *testing.T) {
	repo := NewTodoRepo(setupTestDB(t))
	newtodos := []*model.Todo{{Title: "First Todo", Complete: false},
		{Title: "Second Todo", Complete: true},
		{Title: "Third Todo", Complete: false},
	}
	for _, todo := range newtodos {
		err := repo.Create(todo)
		if err != nil {
			t.Fatalf("failed to create first todo: %v", err)
		}
	}

	todos, err := repo.GetAll()
	if err != nil {
		t.Fatalf("failed to get all todos: %v", err)
	}

	for _, todo := range newtodos {
		ok := slices.Contains(todos, *todo)
		if !ok{
			t.Errorf("expected todo %+v to be in the list, but it was not found", *todo)
		}
	}
}


func TestDelete(t *testing.T) {
	repo := NewTodoRepo(setupTestDB(t))
	todo := &model.Todo{Title: "Todo to be deleted", Complete: false}
	err := repo.Create(todo)
	if err != nil {
		t.Fatalf("failed to create todo: %v", err)
	}

	testCases := []struct {
		name    string
		id      int64
		wantErr bool
	}{
		{
			name:    "valid id",
			id:      todo.Id,
			wantErr: false,
		},
		{
			name:    "invalid id",
			id:      -98,
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.Delete(tc.id)
			if tc.wantErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}