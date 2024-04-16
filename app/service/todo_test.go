package service

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/swarajbhagat11/todo-app/app/repository"
	"github.com/swarajbhagat11/todo-app/err"
	"github.com/swarajbhagat11/todo-app/models"
)

func TestGetByIndexShouldSuccess(t *testing.T) {
	// mock intialization
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository.NewMockTodo(ctrl)
	serv := NewTodoService(mockRepo)
	mockTodo := models.Todo{Id: 1, UserId: 1, Title: "test", Completed: true}

	// assertion
	mockRepo.EXPECT().GetByIndex(1).Return(mockTodo, nil).Times(1)

	// call function
	todo, err := serv.GetByIndex(1)

	// assertion
	assert.Equal(t, mockTodo, todo)
	assert.Nil(t, err)
}

func TestReadDataShouldSuccess(t *testing.T) {
	// mock intialization
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository.NewMockTodo(ctrl)
	serv := NewTodoService(mockRepo)
	httpGet = func(url string) (resp *http.Response, err error) {
		return &http.Response{Body: io.NopCloser(strings.NewReader(`{"userId": 1, "id": 1, "title": "Sample TODO", "completed": false}`))}, nil
	}

	// assertion
	mockRepo.EXPECT().AddAllTodo(gomock.Any()).Return(true, nil).Times(1)

	// call function
	serv.ReadData()
}

func TestReadDataShouldFailWhenAddAllTodoFailure(t *testing.T) {
	// mock intialization
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository.NewMockTodo(ctrl)
	serv := NewTodoService(mockRepo)
	httpGet = func(url string) (resp *http.Response, err error) {
		return &http.Response{Body: io.NopCloser(strings.NewReader(`{"userId": 1, "id": 1, "title": "Sample TODO", "completed": false}`))}, nil
	}
	appError := err.NewAppError(500, "Failed")

	// assertion
	mockRepo.EXPECT().AddAllTodo(gomock.Any()).Return(false, &appError).Times(1)

	// call function
	serv.ReadData()
}

func TestFetchTodoShouldSuccess(t *testing.T) {
	// mock intialization
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"userId": 1, "id": 1, "title": "Sample TODO", "completed": false}`))
	}))
	defer mockServer.Close()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository.NewMockTodo(ctrl)
	serv := NewTodoService(mockRepo)

	todoChan := make(chan models.Todo)

	var wg sync.WaitGroup
	wg.Add(1)

	// call function
	go serv.fetchTodo(mockServer.URL, todoChan, &wg)

	go func() {
		wg.Wait()
		close(todoChan)
	}()
	todo := <-todoChan

	// assertion
	assert.Equal(t, "Sample TODO", todo.Title)
	assert.Equal(t, false, todo.Completed)
}

func TestFetchTodoShouldFailWhenHttpGetHasError(t *testing.T) {
	// mock intialization
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository.NewMockTodo(ctrl)
	serv := NewTodoService(mockRepo)

	todoChan := make(chan models.Todo)

	var wg sync.WaitGroup
	wg.Add(1)

	httpGet = func(string) (resp *http.Response, err error) {
		return nil, errors.New("failed")
	}

	// call function
	go serv.fetchTodo("http://localhost:8080/todo", todoChan, &wg)

	go func() {
		wg.Wait()
		close(todoChan)
	}()
	<-todoChan

	// assertion
	assert.Equal(t, 0, len(todoChan))
}

func TestFetchTodoShouldFailWhenDecodeTodoHasError(t *testing.T) {
	// mock intialization
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(nil))
	}))
	defer mockServer.Close()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := repository.NewMockTodo(ctrl)
	serv := NewTodoService(mockRepo)

	todoChan := make(chan models.Todo)

	var wg sync.WaitGroup
	wg.Add(1)

	// call function
	go serv.fetchTodo(mockServer.URL, todoChan, &wg)

	go func() {
		wg.Wait()
		close(todoChan)
	}()
	<-todoChan

	// assertion
	assert.Equal(t, 0, len(todoChan))
}
