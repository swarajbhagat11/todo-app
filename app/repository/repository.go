package repository

import (
	"github.com/swarajbhagat11/todo-app/err"
	"github.com/swarajbhagat11/todo-app/models"
)

type Todo interface {
	AddAllTodo(todoListToAdd []models.Todo) (bool, *err.AppError)
	GetByIndex(index int) (models.Todo, *err.AppError)
}

type Repository struct {
	Todo
}

func NewRepository() *Repository {
	return &Repository{
		Todo: NewTodoRepository(),
	}
}
