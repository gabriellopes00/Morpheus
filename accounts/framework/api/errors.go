package api

import (
	"errors"
	"net/http"
)

var (
	ErrInternalServer = errors.New("unexpected internal server error")
	ErrInvalidRequest = errors.New("invalid request")
)

type ApiErr interface {
	Message() string
	Status() int
	Name() string
}

type httpApiErr struct {
	ErrMessage string `json:"message,omitempty"`
	ErrStatus  int    `json:"status,omitempty"`
	ErrName    string `json:"name,omitempty"`
}

func (e *httpApiErr) Name() string {
	return e.ErrName
}

func (e *httpApiErr) Message() string {
	return e.ErrMessage
}

func (e *httpApiErr) Status() int {
	return e.ErrStatus
}

func NewUnauthorizedError(message string) ApiErr {
	return &httpApiErr{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrName:    "unauthorized_error",
	}
}

func NewForbiddenError(message string) ApiErr {
	return &httpApiErr{
		ErrMessage: message,
		ErrStatus:  http.StatusForbidden,
		ErrName:    "forbidden_error",
	}
}

func NewUnprocessableEntityError(message string) ApiErr {
	return &httpApiErr{
		ErrMessage: message,
		ErrStatus:  http.StatusUnprocessableEntity,
		ErrName:    "unprocessable_entity",
	}
}

func NewInternalServerError(message string) ApiErr {
	return &httpApiErr{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrName:    "server_error",
	}
}
