package repositories

import (
	"database/sql"
	"events/domain/entities"
)

type AccountRepository interface {
	Create(account *entities.Account) error
	Delete(accountId string) error
}

type pgAccountRepository struct {
	Db *sql.DB
}

func NewPgAccountRepository(connection *sql.DB) *pgAccountRepository {
	return &pgAccountRepository{connection}
}

func (repo *pgAccountRepository) Create(account *entities.Account) error {
	stm, err := repo.Db.Prepare("INSERT INTO accounts VALUES ($1)")
	if err != nil {
		return err
	}

	defer stm.Close()

	_, err = stm.Exec(account.Id)

	return err
}

func (repo *pgAccountRepository) Delete(accountId string) error {
	stm, err := repo.Db.Prepare("DELETE FROM accounts WHERE id = $1")
	if err != nil {
		return err
	}

	defer stm.Close()

	_, err = stm.Exec(accountId)

	return err
}
