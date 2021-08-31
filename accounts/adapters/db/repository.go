package db

import (
	"accounts/domain"
	"database/sql"
	"log"
)

type pgAccountRepository struct {
	Db *sql.DB
}

func NewPgAccountRepository(db *sql.DB) *pgAccountRepository {
	return &pgAccountRepository{
		Db: db,
	}
}

func (r *pgAccountRepository) Create(account *domain.Account) error {
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

func (r *pgAccountRepository) FindByEmail(email string) (*domain.Account, error) {
	stm, err := r.Db.Prepare("SELECT * FROM accounts WHERE email=$1")
	if err != nil {
		return nil, err
	}

	defer stm.Close()

	var account domain.Account

	err = stm.QueryRow(email).Scan(
		&account.Id,
		&account.Name,
		&account.Email,
		&account.Password,
		&account.AvatarUrl,
		&account.CreatedAt,
	)
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	return &account, nil

}