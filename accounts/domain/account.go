package domain

import (
	"errors"
	"fmt"
	"time"

	gouuid "github.com/satori/go.uuid"
)

// Account entity
type Account struct {
	Id        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	AvatarUrl string    `json:"avatar_url,omitempty"`
	BirthDate time.Time `json:"birth_date,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func NewAccount(name, email, password, avatarUrl, birthDate string) (*Account, error) {
	account := &Account{
		Id:        gouuid.NewV4().String(),
		Name:      name,
		Email:     email,
		Password:  password,
		AvatarUrl: avatarUrl,
		CreatedAt: time.Now().Local(),
	}

	parsed, err := time.Parse(time.RFC3339, birthDate)
	if err != nil {
		fmt.Println("here")
		return nil, errors.New("invalid account birth date format")
	}
	account.BirthDate = parsed

	return account, nil
}

// Account usecases

type CreateAccountDTO struct {
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	RG        string `json:"rg,omitempty"`
	BirthDate string `json:"birth_date,omitempty"`
}

type CreateAccount interface {
	Create(data *CreateAccountDTO) (*Account, error)
}

type AuthModel struct {
	AccessToken  string
	RefreshToken string
}

type AuthAccount interface {
	Auth(email, password string) (*AuthModel, error)
}

type RefreshAuth interface {
	Refresh(refreshToken string) (*AuthModel, error)
}

type DeleteAccount interface {
	Delete(accountId string) error
}

type UpdateAccountDTO struct {
	Name      string `json:"name,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	RG        string `json:"rg,omitempty"`
	BirthDate string `json:"birth_date,omitempty"`
}
type UpdateAccount interface {
	Update(accountId string, data *UpdateAccountDTO) (*Account, error)
}

type GetAccount interface {
	GetById(accountId string) (*Account, error)
}
