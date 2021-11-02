package db

import (
	"accounts/domain/entities"
	"database/sql"
	"errors"
)

type pgAccountRepository struct {
	Db *sql.DB
}

func NewPgAccountRepository(db *sql.DB) *pgAccountRepository {
	return &pgAccountRepository{
		Db: db,
	}
}

func (r *pgAccountRepository) Create(account *entities.Account) error {
	stm, err := r.Db.Prepare(`
		INSERT INTO accounts (id,
							 name,
							 email,
							 password,
							 avatar_url,
							 birth_date,
							 created_at,
							 updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
	`)
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
		account.BirthDate,
		account.CreatedAt,
		account.UpdatedAt)

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

func (r *pgAccountRepository) FindByEmail(email string) (*entities.Account, error) {
	stm, err := r.Db.Prepare(`
		SELECT id,
			   name,
			   email,
			   password,
			   avatar_url,
			   birth_date,
			   created_at,
			   updated_at
		FROM accounts WHERE email = $1`)
	if err != nil {
		return nil, err
	}

	defer stm.Close()

	var account entities.Account

	err = stm.QueryRow(email).Scan(
		&account.Id,
		&account.Name,
		&account.Email,
		&account.Password,
		&account.AvatarUrl,
		&account.BirthDate,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &account, nil

}

func (r *pgAccountRepository) FindById(id string) (*entities.Account, error) {
	stm, err := r.Db.Prepare(`
		SELECT id,
			   name,
			   email,
			   password,
			   avatar_url,
			   birth_date,
			   created_at,
			   updated_at
		FROM accounts WHERE id = $1`)
	if err != nil {
		return nil, err
	}

	defer stm.Close()

	var account entities.Account

	err = stm.QueryRow(id).Scan(
		&account.Id,
		&account.Name,
		&account.Email,
		&account.Password,
		&account.AvatarUrl,
		&account.BirthDate,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &account, nil

}

func (r *pgAccountRepository) ExistsId(id string) (bool, error) {
	stm, err := r.Db.Prepare("SELECT id FROM accounts WHERE id=$1")
	if err != nil {
		return false, err
	}

	defer stm.Close()

	var accountId string

	err = stm.QueryRow(id).Scan(&accountId)
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (r *pgAccountRepository) Delete(id string) error {
	stm, err := r.Db.Prepare("DELETE FROM accounts WHERE id = $1;")
	if err != nil {
		return err
	}

	defer stm.Close()

	_, err = stm.Exec(id)
	return err

}

func (r *pgAccountRepository) Update(account *entities.Account) error {
	stm, err := r.Db.Prepare(`
		UPDATE accounts
		SET name = $1,
			avatar_url = $2,
			birth_date = $3,
			updated_at = $4
		WHERE id = $5;
	`)
	if err != nil {
		return err
	}

	defer stm.Close()

	_, err = stm.Exec(
		account.Name,
		account.AvatarUrl,
		account.BirthDate,
		account.UpdatedAt,
		account.Id,
	)

	return err
}
