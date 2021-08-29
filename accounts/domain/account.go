package domain

import (
	"errors"
	"time"
)

// Account entity
type Account struct {
	Id        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	AvatarUrl string    `json:"avatar_url,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

var (
	ErrEmailAlreadyInUse = errors.New("email already in use")
)

// Create account usecase
type CreateAccount interface {
	Create(data Account) (*Account, error)
}
