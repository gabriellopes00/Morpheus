package application

import (
	"accounts/pkg/db"
	"errors"
)

type DeleteAccount struct {
	Repository db.Repository
}

func NewDeleteAccount(Repository db.Repository) *DeleteAccount {
	return &DeleteAccount{
		Repository: Repository,
	}
}

var (
	ErrIdNotFound = errors.New("id not found")
)

func (u *DeleteAccount) Delete(accountId string) error {
	existingAccount, err := u.Repository.ExistsId(accountId)
	if err != nil {
		return err
	}

	if !existingAccount {
		return ErrIdNotFound
	}

	err = u.Repository.Delete(accountId)
	if err != nil {
		return err
	}

	return nil
}
