package application

import (
	"events/framework/db/repositories"
	"log"
)

type DeleteAccount struct {
	repository repositories.AccountRepository
}

func NewDeleteAccount(repo repositories.AccountRepository) *DeleteAccount {
	return &DeleteAccount{
		repository: repo,
	}
}

func (c DeleteAccount) Delete(accountId string) error {
	log.Println(accountId)
	return c.repository.Delete(accountId)
}
