package application

import (
	"events/domain/entities"
	"events/framework/db/repositories"
)

type createAccount struct {
	repository repositories.AccountRepository
}

func NewCreateAccount(repo repositories.AccountRepository) *createAccount {
	return &createAccount{
		repository: repo,
	}
}

func (c createAccount) Create(account *entities.Account) error {
	return c.repository.Create(account)
}
