package application

import (
	"accounts/domain/entities"
	"accounts/pkg/auth"
	"accounts/pkg/db"
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
	Repository   db.Repository
	AuthProvider auth.AuthProvider
}

func NewCreateAccount(Repository db.Repository, AuthProvider auth.AuthProvider) *CreateAccount {
	return &CreateAccount{
		Repository:   Repository,
		AuthProvider: AuthProvider,
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

	// err = c.Repository.Create(account)
	// if err != nil {
	// 	return nil, err
	// }

	err = c.AuthProvider.CreateUser(
		auth.AuthProviderUser{
			Id:        account.Id,
			Name:      account.Name,
			Email:     account.Email,
			Password:  account.Password,
			CreatedAt: account.CreatedAt,
		},
	)

	return account, nil
}
