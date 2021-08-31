package interfaces

import "accounts/domain"

type Repository interface {
	Create(account *domain.Account) error
	Exists(email string) (bool, error)
	FindByEmail(email string) (*domain.Account, error)
}
