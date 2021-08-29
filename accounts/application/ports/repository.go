package ports

import "accounts/domain"

type Repository interface {
	Create(account *domain.Account) error
	Exists(email string) (bool, error)
}
