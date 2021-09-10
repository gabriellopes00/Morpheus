package api

import (
	"net/http"
)

type ApiErr interface {
	Message() string
	Status() int
	Error() string
}

type apiErr struct {
	ErrMessage string `json:"message"`
	ErrStatus  int    `json:"status"`
	Err        string `json:"error"`
}

func (e *apiErr) Error() string {
	return e.Err
}

func (e *apiErr) Message() string {
	return e.ErrMessage
}

func (e *apiErr) Status() int {
	return e.ErrStatus
}

func NewNotFoundError(message string) ApiErr {
	return &apiErr{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		Err:        "not_found",
	}
}

func NewBadRequestError(message string) ApiErr {
	return &apiErr{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		Err:        "bad_request",
	}
}
func NewUnprocessableEntityError(message string) ApiErr {
	return &apiErr{
		ErrMessage: message,
		ErrStatus:  http.StatusUnprocessableEntity,
		Err:        "invalid_request",
	}
}

func NewInternalServerError(message string) ApiErr {
	return &apiErr{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		Err:        "server_error",
	}
}
