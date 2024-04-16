package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/swarajbhagat11/todo-app/app/service"
)

type Handler struct {
	Todo *TodoHandler
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		Todo: NewTodoHandler(services.Todo),
	}
}

func (h *Handler) InitRoutes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", h.Todo.Ping)
	r.Get("/todos/{index}", h.Todo.GetByIndex)

	return r
}
