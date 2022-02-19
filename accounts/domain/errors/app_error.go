package app_error

import "fmt"

type AppError struct {
	error
	Type    string `json:"type,omitempty"`
	Details string `json:"details,omitempty"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.Type, e.Details)
}

func NewAppError(errType, details string) *AppError {
	return &AppError{
		Type:    errType,
		Details: details,
	}

}

func IsAppError(error interface{}) bool {
	if _, ok := error.(AppError); ok {
		return true
	} else if _, ok := error.(*AppError); ok {
		return true
	}

	return false
}
