package db

import (
	"accounts/domain/entities"
	"accounts/domain/usecases"
	"database/sql"
	"fmt"
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
							 created_at,
							 updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7);
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
	stm, err := r.Db.Prepare("SELECT * FROM accounts WHERE email=$1")
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
		&account.CreatedAt,
	)
	if err != nil {
		return nil, nil
	}

	return &account, nil

}

func (r *pgAccountRepository) FindById(id string) (*entities.Account, error) {
	stm, err := r.Db.Prepare("SELECT * FROM accounts WHERE id=$1")
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
		&account.CreatedAt,
	)
	if err != nil {
		return nil, nil
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
	stm, err := r.Db.Prepare("DELETE FROM accounts WHERE id=$1")
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer stm.Close()

	_, err = stm.Exec(id)
	if err != nil {
		fmt.Println(err)

		return err
	}

	return nil

}

func (r *pgAccountRepository) Update(accountId string, data *usecases.UpdateAccountDTO) (*entities.Account, error) {
	stm, err := r.Db.Prepare(`
		UPDATE accounts
		SET name = $1,
			avatar_url = $2,
			rg = $3,
			birth_date = $4
		WHERE id = $5 
		RETURNING id,
				  name,
				  email,
				  password,
				  avatar_url,
				  birth_date,
				  created_at;
	`)
	if err != nil {
		return nil, err
	}

	defer stm.Close()

	account := entities.Account{}

	row := stm.QueryRow(
		accountId,
		data.Name,
		data.RG,
		data.AvatarUrl,
		data.BirthDate,
	)

	err = row.Scan(
		&account.Id,
		&account.Name,
		&account.Email,
		&account.Password,
		&account.AvatarUrl,
		&account.BirthDate,
		&account.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &account, err
}
