package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/primekobie/todos/model"
)

type Handler struct {
	store model.TodoRepository
}

func NewHandler(store model.TodoRepository) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var data model.Todo
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.store.Create(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(data)
		return
	}
}

func (h *Handler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var data model.Todo
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.store.Update(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
		return
	}
}

func (h *Handler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.store.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todos)
		return
	}
}

func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.store.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		return
	}
}
