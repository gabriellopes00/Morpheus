package application

import (
	"events/domain/entities"
	"events/framework/db/repositories"
)

type CreateAccount struct {
	repository repositories.AccountRepository
}

func NewCreateAccount(repo repositories.AccountRepository) *CreateAccount {
	return &CreateAccount{
		repository: repo,
	}
}

func (c CreateAccount) Create(account *entities.Account) error {
	return c.repository.Create(account)
}
