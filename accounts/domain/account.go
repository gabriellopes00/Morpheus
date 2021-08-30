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

type AuthCredentials struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

var (
	ErrEmailAlreadyInUse = errors.New("email already in use")
	ErrUnregisteredEmail = errors.New("unregistered email")
)

// Account usecases
type AccountUsecase interface {
	Create(data Account) (*Account, error)
	Auth(data AuthCredentials) (string, error)
}
