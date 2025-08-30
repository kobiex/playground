package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/primekobie/todos/model"
	"github.com/primekobie/todos/model/mock"
	"go.uber.org/mock/gomock"
)

func TestCreateTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		body       interface{}
		mockErr    error
		expectCode int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				body:       model.Todo{Id: 1, Title: "Test", Complete: false},
				mockErr:    nil,
				expectCode: http.StatusCreated,
			},
		},
		{
			name: "bad request",
			args: args{
				body:       "{invalid json",
				mockErr:    nil,
				expectCode: http.StatusBadRequest,
			},
		},
		{
			name: "internal error",
			args: args{
				body:       model.Todo{Id: 2, Title: "Err", Complete: false},
				mockErr:    errors.New("db error"),
				expectCode: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mock.NewMockTodoRepository(ctrl)
			handler := NewHandler(mockRepo)

			var req *http.Request
			switch v := tt.args.body.(type) {
			case model.Todo:
				b, _ := json.Marshal(v)
				req = httptest.NewRequest("POST", "/todos", bytes.NewReader(b))
				mockRepo.EXPECT().Create(&v).Return(tt.args.mockErr).AnyTimes()
			case string:
				req = httptest.NewRequest("POST", "/todos", bytes.NewReader([]byte(v)))
			}

			w := httptest.NewRecorder()
			handler.CreateTodo(w, req)
			if w.Code != tt.args.expectCode {
				t.Errorf("expected %d, got %d", tt.args.expectCode, w.Code)
			}
		})
	}
}

func TestUpdateTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		body       interface{}
		mockErr    error
		expectCode int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				body:       model.Todo{Id: 1, Title: "Updated", Complete: true},
				mockErr:    nil,
				expectCode: http.StatusOK,
			},
		},
		{
			name: "bad request",
			args: args{
				body:       "{invalid json",
				mockErr:    nil,
				expectCode: http.StatusBadRequest,
			},
		},
		{
			name: "internal error",
			args: args{
				body:       model.Todo{Id: 2, Title: "Err", Complete: false},
				mockErr:    errors.New("db error"),
				expectCode: http.StatusInternalServerError,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mock.NewMockTodoRepository(ctrl)
			handler := NewHandler(mockRepo)

			var req *http.Request
			switch v := tt.args.body.(type) {
			case model.Todo:
				b, _ := json.Marshal(v)
				req = httptest.NewRequest("PUT", "/todos", bytes.NewReader(b))
				mockRepo.EXPECT().Update(&v).Return(tt.args.mockErr).AnyTimes()
			case string:
				req = httptest.NewRequest("PUT", "/todos", bytes.NewReader([]byte(v)))
			}

			w := httptest.NewRecorder()
			handler.UpdateTodo(w, req)
			if w.Code != tt.args.expectCode {
				t.Errorf("expected %d, got %d", tt.args.expectCode, w.Code)
			}
		})
	}
}

func TestGetAllTodos(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name       string
		mockTodos  []model.Todo
		mockErr    error
		expectCode int
	}{
		{
			name:       "success",
			mockTodos:  []model.Todo{{Id: 1, Title: "A", Complete: false}},
			mockErr:    nil,
			expectCode: http.StatusOK,
		},
		{
			name:       "internal error",
			mockTodos:  nil,
			mockErr:    errors.New("db error"),
			expectCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mock.NewMockTodoRepository(ctrl)
			handler := NewHandler(mockRepo)

			mockRepo.EXPECT().GetAll().Return(tt.mockTodos, tt.mockErr).AnyTimes()

			req := httptest.NewRequest("GET", "/todos", nil)
			w := httptest.NewRecorder()
			handler.GetAllTodos(w, req)
			if w.Code != tt.expectCode {
				t.Errorf("expected %d, got %d", tt.expectCode, w.Code)
			}
		})
	}
}

func TestDeleteTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name       string
		id         string
		mockErr    error
		expectCode int
	}{
		{
			name:       "success",
			id:         "1",
			mockErr:    nil,
			expectCode: http.StatusOK,
		},
		{
			name:       "bad request",
			id:         "abc",
			mockErr:    nil,
			expectCode: http.StatusBadRequest,
		},
		{
			name:       "internal error",
			id:         "2",
			mockErr:    errors.New("db error"),
			expectCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mock.NewMockTodoRepository(ctrl)
			handler := NewHandler(mockRepo)

			req := httptest.NewRequest("DELETE", "/todos?id="+tt.id, nil)
			w := httptest.NewRecorder()

			if idInt, err := strconv.ParseInt(tt.id, 10, 64); err == nil {
				mockRepo.EXPECT().Delete(idInt).Return(tt.mockErr).AnyTimes()
			}

			handler.DeleteTodo(w, req)
			if w.Code != tt.expectCode {
				t.Errorf("expected %d, got %d", tt.expectCode, w.Code)
			}
		})
	}
}
