package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"github.com/swarajbhagat11/todo-app/app/service"
	"github.com/swarajbhagat11/todo-app/helpers"
)

type TodoHandler struct {
	ser service.Todo
}

func NewTodoHandler(ser service.Todo) *TodoHandler {
	return &TodoHandler{ser: ser}
}

func (h *TodoHandler) Ping(w http.ResponseWriter, r *http.Request) {
	log.Println("[TodoHandler => Ping] Started execution...")
	w.WriteHeader(http.StatusOK)
	log.Println("[TodoHandler => Ping] Execution completed...")
	w.Write([]byte("Todo API is live!"))
}

func (h *TodoHandler) GetByIndex(w http.ResponseWriter, r *http.Request) {
	log.Println("[TodoHandler => GetByIndex] Started execution...")
	index := chi.URLParam(r, "index")
	indexVal, err := strconv.Atoi(index)
	if err != nil {
		log.Errorln("[TodoHandler => GetByIndex] Error while index conversion.", err)
		helpers.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	todo, er := h.ser.GetByIndex(indexVal)
	if er != nil {
		log.Errorln("[TodoHandler => GetByIndex] Error at service GetByIndex call.", err)
		helpers.RespondWithError(w, er, er.StatusCode)
		return
	}

	log.Println("[TodoHandler => GetByIndex] Execution completed...")
	helpers.WriteJSON(w, todo, http.StatusOK)
}
