package application

import (
	"events/framework/db/repositories"
	"log"
)

type deleteAccount struct {
	repository repositories.AccountRepository
}

func NewDeleteAccount(repo repositories.AccountRepository) *deleteAccount {
	return &deleteAccount{
		repository: repo,
	}
}

func (c deleteAccount) Delete(accountId string) error {
	log.Println(accountId)
	return c.repository.Delete(accountId)
}
