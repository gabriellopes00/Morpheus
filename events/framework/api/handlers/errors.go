package handlers

import "errors"

var (
	ErrBadRequest          = errors.New("invalid request")
	ErrUnprocessableEntity = errors.New("unprocessable entity")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrInternalServer      = errors.New("unexpected internal server error")
)
