package api

import (
	"net/http"
)

type HttpApiErr struct {
	ErrMessage string `json:"message,omitempty"`
	ErrStatus  int    `json:"status,omitempty"`
	ErrName    string `json:"name,omitempty"`
}

func NewUnauthorizedError(message string) *HttpApiErr {
	return &HttpApiErr{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrName:    "unauthorized_error",
	}
}

func NewForbiddenError(message string) *HttpApiErr {
	return &HttpApiErr{
		ErrMessage: message,
		ErrStatus:  http.StatusForbidden,
		ErrName:    "forbidden_error",
	}
}

func NewUnprocessableEntityError(message string) *HttpApiErr {
	return &HttpApiErr{
		ErrMessage: message,
		ErrStatus:  http.StatusUnprocessableEntity,
		ErrName:    "unprocessable_entity",
	}
}

func NewInternalServerError(message string) *HttpApiErr {
	return &HttpApiErr{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrName:    "server_error",
	}
}
