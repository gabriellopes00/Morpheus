package usecases

import "accounts/domain/entities"

type CreateAccount interface {
	Create(data entities.Account) (*entities.Account, error)
}
