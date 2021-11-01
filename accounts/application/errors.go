package application

import "errors"

var (
	ErrEmailAlreadyInUse = errors.New("email already in use")
)
