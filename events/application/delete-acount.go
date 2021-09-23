package application

import (
	"events/framework/db/repositories"
	"log"
)

type DeleteAccountUsecase struct {
	Repository repositories.AccountRepository
}

func (c DeleteAccountUsecase) Delete(accountId string) error {
	log.Println(accountId)
	return c.Repository.Delete(accountId)
}
