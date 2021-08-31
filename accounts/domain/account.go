package domain

import "time"

// Account entity
type Account struct {
	Id        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	AvatarUrl string    `json:"avatar_url,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Account usecases

type CreateAccount interface {
	Create(data Account) (*Account, error)
}

type AuthAccount interface {
	Auth(email, password string) (string, error)
}

type DeleteAccount interface {
	Delete(accountId string) error
}

type UpdateAccount interface {
	Update(accountId string, data Account) (*Account, error)
}

type GetAccount interface {
	GetById(accountId string) (*Account, error)
}
