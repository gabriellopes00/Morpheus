package handlers

import "errors"

var (
	ErrInvalidRequest      = errors.New("invalid request")
	ErrUnprocessableEntity = errors.New("unprocessable entity")
	ErrUnauthorized        = errors.New("unauthorized event creation")
	ErrInternalServer      = errors.New("unexpected internal server error")
)
