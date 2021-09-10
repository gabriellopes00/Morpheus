package usecases

import "events/domain/entities"

type CreateAccount interface {
	Create(account *entities.Account) error
}

type DeleteAccount interface {
	Delete(accountId string) error
}
