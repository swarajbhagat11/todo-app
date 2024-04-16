package helpers

import (
	"encoding/json"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/swarajbhagat11/todo-app/models"
)

type ErrResponseJSON struct {
	ErrMsg string `json:"error"`
}

func RespondWithError(w http.ResponseWriter, receivedError error, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	// Init new error struct
	errMsg := ErrResponseJSON{ErrMsg: receivedError.Error()}
	err := json.NewEncoder(w).Encode(errMsg)
	if err != nil {
		log.Errorln("[Helpers => RespondWithError] error encoding json:", err)
	}
}

func DecodeTodoJSON(body io.ReadCloser) (models.Todo, error) {
	todo := models.Todo{}

	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&todo); err != nil {
		log.Errorln("[Helpers => DecodeTodoJSON] error decoding json:", err)
		return todo, err
	}

	return todo, nil
}

func WriteJSON(w http.ResponseWriter, data any, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Errorln("[Helpers => WriteJSON] error encoding json:", err)
		return err
	}
	return nil
}
