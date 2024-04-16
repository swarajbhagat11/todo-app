package err

import (
	"errors"
	"fmt"
)

type AppError struct {
	Err        error
	StatusCode int
}

func NewAppError(statusCode int, message string) AppError {
	e := AppError{}
	e.StatusCode = 404
	e.Err = errors.New(message)

	return e
}

func (r *AppError) Error() string {
	return fmt.Sprintf("%v", r.Err)
}

var NotFoundError = NewAppError(404, "resource not found")
