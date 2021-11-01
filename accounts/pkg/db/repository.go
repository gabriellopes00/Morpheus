package db

import (
	"accounts/domain"
	"accounts/domain/entities"
)

type Repository interface {
	Create(account *entities.Account) error
	Exists(email string) (bool, error)
	ExistsId(param string) (bool, error)
	FindByEmail(email string) (*domain.Account, error)
	FindById(id string) (*domain.Account, error)
	Delete(id string) error
	Update(accountId string, data *domain.UpdateAccountDTO) (*domain.Account, error)
}
