// Code generated by MockGen. DO NOT EDIT.
// Source: app/repository/repository.go

// Package repository is a generated GoMock package.
package repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	err "github.com/swarajbhagat11/todo-app/err"
	models "github.com/swarajbhagat11/todo-app/models"
)

// MockTodo is a mock of Todo interface.
type MockTodo struct {
	ctrl     *gomock.Controller
	recorder *MockTodoMockRecorder
}

// MockTodoMockRecorder is the mock recorder for MockTodo.
type MockTodoMockRecorder struct {
	mock *MockTodo
}

// NewMockTodo creates a new mock instance.
func NewMockTodo(ctrl *gomock.Controller) *MockTodo {
	mock := &MockTodo{ctrl: ctrl}
	mock.recorder = &MockTodoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodo) EXPECT() *MockTodoMockRecorder {
	return m.recorder
}

// AddAllTodo mocks base method.
func (m *MockTodo) AddAllTodo(todoListToAdd []models.Todo) (bool, *err.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAllTodo", todoListToAdd)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*err.AppError)
	return ret0, ret1
}

// AddAllTodo indicates an expected call of AddAllTodo.
func (mr *MockTodoMockRecorder) AddAllTodo(todoListToAdd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAllTodo", reflect.TypeOf((*MockTodo)(nil).AddAllTodo), todoListToAdd)
}

// GetByIndex mocks base method.
func (m *MockTodo) GetByIndex(index int) (models.Todo, *err.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIndex", index)
	ret0, _ := ret[0].(models.Todo)
	ret1, _ := ret[1].(*err.AppError)
	return ret0, ret1
}

// GetByIndex indicates an expected call of GetByIndex.
func (mr *MockTodoMockRecorder) GetByIndex(index interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIndex", reflect.TypeOf((*MockTodo)(nil).GetByIndex), index)
}
