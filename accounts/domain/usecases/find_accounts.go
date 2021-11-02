package usecases

import "accounts/domain/entities"

type FindAccount interface {
	FindById(accountId string) (*entities.Account, error)
}
