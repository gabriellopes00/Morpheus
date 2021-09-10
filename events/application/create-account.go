package application

import (
	"events/domain/entities"
	"events/framework/db/repositories"
)

type CreateAccountUsecase struct {
	Repository repositories.AccountRepository
}

func (c CreateAccountUsecase) Create(account *entities.Account) error {
	return c.Repository.Create(account)
}
