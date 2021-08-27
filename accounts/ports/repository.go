package ports

import "accounts/domain/entities"

type Repository interface {
	Create(account *entities.Account) error
	Exists(email string) (bool, error)
}
