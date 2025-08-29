package main

import (
	"net/http/httptest"
	"testing"

	"github.com/bitkobie/todos/model"
	"github.com/bitkobie/todos/model/mock"
	"go.uber.org/mock/gomock"
)

func TestCreateTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockTodoRepository(ctrl)
	_ = NewHandler(mockRepo)

	todo := &model.Todo{Title: "Test Todo", Complete: false}
	mockRepo.EXPECT().Create(todo).Return(nil)

	_ = httptest.NewRecorder()
	_ = httptest.NewRequest("POST", "/todos", nil)

}
