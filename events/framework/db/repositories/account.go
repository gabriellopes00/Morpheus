package repositories

import (
	"events/domain/entities"

	"gorm.io/gorm"
)

type AccountRepository interface {
	Create(account *entities.Account) error
	Delete(accountId string) error
}

type pgAccountRepository struct {
	Db *gorm.DB
}

func NewPgAccountRepository(connection *gorm.DB) *pgAccountRepository {
	return &pgAccountRepository{connection}
}

func (repo *pgAccountRepository) Create(account *entities.Account) error {
	err := repo.Db.Create(&account).Error
	return err
}

func (repo *pgAccountRepository) Delete(accountId string) error {
	err := repo.Db.Where("id = ?", accountId).Delete(&entities.Account{}).Error
	return err
}
