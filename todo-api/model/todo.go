package model

type Todo struct {
	Id       int64  `json:"id"`
	Title    string `json:"title"`
	Complete bool   `json:"complete"`
}

type TodoRepository interface {
	Create(todo *Todo) error
	Update(todo *Todo) error
	GetAll() ([]Todo, error)
	Delete(id int64) error
}
