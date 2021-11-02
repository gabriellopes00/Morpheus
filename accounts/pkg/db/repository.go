package db

import (
	"accounts/domain/entities"
)

type Repository interface {
	Create(account *entities.Account) error
	Exists(email string) (bool, error)
	ExistsId(param string) (bool, error)
	FindByEmail(email string) (*entities.Account, error)
	FindById(id string) (*entities.Account, error)
	Delete(id string) error
	Update(data *entities.Account) error
}
