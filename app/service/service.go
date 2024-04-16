package service

import (
	"github.com/swarajbhagat11/todo-app/app/repository"
	"github.com/swarajbhagat11/todo-app/err"
	"github.com/swarajbhagat11/todo-app/models"
)

type Todo interface {
	GetByIndex(index int) (models.Todo, *err.AppError)
	ReadData()
}

type Service struct {
	Todo
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Todo: NewTodoService(repo.Todo),
	}
}
