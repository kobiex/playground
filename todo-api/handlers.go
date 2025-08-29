package main

import (
	"net/http"

	"github.com/bitkobie/todos/model"
)

type Handler struct{
	store model.TodoRepository
}

func NewHandler(store model.TodoRepository) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request){
	
}

func (h *Handler) UpdateTodo(w http.ResponseWriter, r *http.Request){
	
}

func (h *Handler) GetAllTodos(w http.ResponseWriter, r *http.Request){
	
}

func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request){
	
}

