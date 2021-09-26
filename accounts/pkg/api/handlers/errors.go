package handlers

import "errors"

var (
	ErrInternalServer = errors.New("unexpected internal server error")
	ErrInvalidRequest = errors.New("invalid request")
)
