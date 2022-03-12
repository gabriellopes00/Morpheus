package application

import (
	app_error "accounts/domain/errors"
)

var (
	ErrEmailAlreadyInUse = app_error.NewAppError("Conflict error", "email already in use")
	ErrUnregisteredEmail = app_error.NewAppError("Not found error", "unregistered email")
)
