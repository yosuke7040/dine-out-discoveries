package response

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrParameterInvalid = errors.New("parameter invalid")

	ErrInvalidInput = errors.New("invalid input")
)

type Error struct {
	statusCode int
	Errors     []string `json:"errors"`
}

func NewError(err error, statusCode int) *Error {
	return &Error{
		statusCode: statusCode,
		Errors:     []string{err.Error()},
	}
}

func NewErrorMessage(messages []string, status int) *Error {
	return &Error{
		statusCode: status,
		Errors:     messages,
	}
}

func (e *Error) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.statusCode)
	return json.NewEncoder(w).Encode(e)
}
