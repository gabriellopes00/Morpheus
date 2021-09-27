package errors

type DomainErr interface {
	Error() string
}

// Validation Error

type ValidationError struct {
	ErrMessage string `json:"message"`
	Field      string `json:"field"`
}

func (e *ValidationError) Error() string {
	return e.ErrMessage
}

func (e *ValidationError) InvalidField() string {
	return e.Field
}

func NewValidationError(message, field string) DomainErr {
	return &ValidationError{
		ErrMessage: message,
		Field:      field,
	}
}

func IsDomainError(err interface{}) bool {
	_, ok := err.(ValidationError)
	return ok
}
