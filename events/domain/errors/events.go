package domain_errors

type DomainErr interface {
	Error() string
}

// Validation Error

type ValidationError struct {
	ErrMessage string      `json:"message,omitempty"`
	Field      string      `json:"field,omitempty"`
	Value      interface{} `json:"value,omitempty"`
}

func (e *ValidationError) Error() string {
	return e.ErrMessage
}

func (e *ValidationError) InvalidField() string {
	return e.Field
}

func NewValidationError(message, field string, value interface{}) DomainErr {
	return &ValidationError{
		ErrMessage: message,
		Field:      field,
		Value:      value,
	}
}

func IsDomainError(err interface{}) bool {
	_, ok := err.(ValidationError)
	return ok
}
