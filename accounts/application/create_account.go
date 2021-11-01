package application

import (
	"accounts/domain/entities"
	"accounts/domain/usecases"
	"accounts/pkg/db"

	"golang.org/x/crypto/bcrypt"
)

type createAccount struct {
	Repository db.Repository
}

func NewCreateAccount(Repository db.Repository) *createAccount {
	return &createAccount{
		Repository: Repository,
	}
}

func (c *createAccount) Create(data *usecases.CreateAccountDTO) (*entities.Account, error) {
	accountExists, err := c.Repository.Exists(data.Email)
	if err != nil {
		return nil, err
	}

	if accountExists {
		return nil, ErrEmailAlreadyInUse
	}

	account, err := entities.NewAccount(data.Name, data.Email, data.Password, data.AvatarUrl, data.BirthDate)
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
