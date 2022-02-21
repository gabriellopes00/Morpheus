package application

import (
	"accounts/domain/entities"
	"accounts/pkg/db"

	"golang.org/x/crypto/bcrypt"
)

type CreateAccountDTO struct {
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	Document  string `json:"document,omitempty"`
	BirthDate string `json:"birth_date,omitempty"`
}

type CreateAccount struct {
	Repository db.Repository
}

func NewCreateAccount(Repository db.Repository) *CreateAccount {
	return &CreateAccount{
		Repository: Repository,
	}
}

func (c *CreateAccount) Create(data *CreateAccountDTO) (*entities.Account, error) {
	accountExists, err := c.Repository.Exists(data.Email)
	if err != nil {
		return nil, err
	}

	if accountExists {
		return nil, ErrEmailAlreadyInUse
	}

	account, err := entities.NewAccount(
		data.Name,
		data.Email,
		data.Password,
		data.AvatarUrl,
		data.BirthDate,
		data.Document,
	)
	if err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	account.Password = string(hash)

	err = c.Repository.Create(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}
