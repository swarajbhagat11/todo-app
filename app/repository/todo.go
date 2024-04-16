package repository

import (
	log "github.com/sirupsen/logrus"
	"github.com/swarajbhagat11/todo-app/err"
	"github.com/swarajbhagat11/todo-app/models"
)

type TodoRepository struct {
	todoList []models.Todo
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{}
}

func (r *TodoRepository) AddAllTodo(todoListToAdd []models.Todo) (bool, *err.AppError) {
	r.todoList = todoListToAdd

	return true, nil
}

func (r *TodoRepository) GetByIndex(index int) (models.Todo, *err.AppError) {
	if index < 1 || len(r.todoList) < index {
		log.Errorln("[TodoRepository => GetByIndex] Todo is out of bound.")
		return models.Todo{}, &err.NotFoundError
	}

	return r.todoList[index-1], nil
}
