package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/swarajbhagat11/todo-app/app/service"
	"github.com/swarajbhagat11/todo-app/err"
	"github.com/swarajbhagat11/todo-app/models"
)

func TestPingShouldSuccess(t *testing.T) {
	// mock intialization
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSer := service.NewMockTodo(ctrl)
	hand := NewTodoHandler(mockSer)

	// init server
	r := chi.NewRouter()
	r.Get("/", hand.Ping)

	// init recorder and request
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	// perform request
	r.ServeHTTP(w, req)

	// assertion
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Todo API is live!", strings.Trim(w.Body.String(), "\n"))
}

func TestGetByIndexShouldSuccess(t *testing.T) {
	// mock intialization
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSer := service.NewMockTodo(ctrl)
	hand := NewTodoHandler(mockSer)
	mockTodo := models.Todo{Id: 1, UserId: 1, Title: "test", Completed: true}
	expectTodo, _ := json.Marshal(mockTodo)
	mockSer.EXPECT().GetByIndex(1).Return(mockTodo, nil).Times(1)

	// init server
	r := chi.NewRouter()
	r.Get("/todo/{index}", hand.GetByIndex)

	// init recorder and request
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/todo/1", nil)

	// perform request
	r.ServeHTTP(w, req)

	// assertion
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(expectTodo), strings.Trim(w.Body.String(), "\n"))
}

func TestGetByIndexShouldFailWhenErrorWhileParsingPathParameter(t *testing.T) {
	// mock intialization
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSer := service.NewMockTodo(ctrl)
	hand := NewTodoHandler(mockSer)
	mockSer.EXPECT().GetByIndex(1).Return(models.Todo{}, nil).Times(0)

	// init server
	r := chi.NewRouter()
	r.Get("/todo/{index}", hand.GetByIndex)

	// init recorder and request
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/todo/wrongValue", nil)

	// perform request
	r.ServeHTTP(w, req)

	// assertion
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetByIndexShouldFailWhenGetByIndexHasError(t *testing.T) {
	// mock intialization
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockSer := service.NewMockTodo(ctrl)
	hand := NewTodoHandler(mockSer)
	errorMsg := "resource not found"
	appError := err.NewAppError(404, errorMsg)
	mockSer.EXPECT().GetByIndex(1).Return(models.Todo{}, &appError).Times(1)

	// init server
	r := chi.NewRouter()
	r.Get("/todo/{index}", hand.GetByIndex)

	// init recorder and request
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/todo/1", nil)

	// perform request
	r.ServeHTTP(w, req)

	// assertion
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, fmt.Sprintf("{\"error\":\"%s\"}", errorMsg), strings.Trim(w.Body.String(), "\n"))
}
