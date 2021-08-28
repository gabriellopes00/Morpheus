package db

import (
	"accounts/domain/entities"
	"database/sql"
)

type pgAccountRepository struct {
	Db *sql.DB
}

func NewPgAccountRepository(db *sql.DB) *pgAccountRepository {
	return &pgAccountRepository{db}
}

func (r *pgAccountRepository) Create(account *entities.Account) error {
	stm, err := r.Db.Prepare("INSERT INTO accounts VALUES ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return err
	}

	defer stm.Close()

	_, err = stm.Exec(
		account.Id,
		account.Name,
		account.Email,
		account.Password,
		account.AvatarUrl,
		account.CreatedAt,
	)

	return err
}

func (r *pgAccountRepository) Exists(email string) (bool, error) {
	stm, err := r.Db.Prepare("SELECT id FROM accounts WHERE email=$1")
	if err != nil {
		return false, err
	}

	defer stm.Close()

	var accountId string

	err = stm.QueryRow(email).Scan(&accountId)
	if err != nil {
		return false, nil
	}

	return true, nil

}
