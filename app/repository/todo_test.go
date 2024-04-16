package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/swarajbhagat11/todo-app/models"
)

func TestAddAllTodoShouldSuccess(t *testing.T) {
	// mock intialization
	repo := NewTodoRepository()
	mockTodo := []models.Todo{{Id: 1, UserId: 1, Title: "test", Completed: true}}
	passedTodo := []models.Todo{{Id: 2, UserId: 2, Title: "test2", Completed: true}, {Id: 3, UserId: 3, Title: "test3", Completed: false}}
	repo.todoList = mockTodo

	// call function
	success, err := repo.AddAllTodo(passedTodo)

	// assertion
	assert.True(t, success)
	assert.Nil(t, err)
	assert.Equal(t, passedTodo, repo.todoList)
}

func TestGetByIndexShouldSuccess(t *testing.T) {
	// mock intialization
	repo := NewTodoRepository()
	mockTodo := models.Todo{Id: 1, UserId: 1, Title: "test", Completed: true}
	repo.todoList = []models.Todo{mockTodo}

	// call function
	todo, err := repo.GetByIndex(1)

	// assertion
	assert.Equal(t, mockTodo, todo)
	assert.Nil(t, err)
}

func TestGetByIndexShouldFailWhenIndexIsOutOfBound(t *testing.T) {
	// mock intialization
	repo := NewTodoRepository()
	mockTodo := models.Todo{Id: 1, UserId: 1, Title: "test", Completed: true}
	repo.todoList = []models.Todo{mockTodo}

	// call function
	todo, err := repo.GetByIndex(2)

	// assertion
	assert.Equal(t, models.Todo{}, todo)
	assert.NotNil(t, err)
}
