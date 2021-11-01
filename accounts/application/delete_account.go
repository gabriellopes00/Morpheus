package application

import (
	"accounts/pkg/db"
	"errors"
)

type deleteUsecase struct {
	Repository db.Repository
}

func NewDeleteUsecase(Repository db.Repository) *deleteUsecase {
	return &deleteUsecase{
		Repository: Repository,
	}
}

var (
	ErrIdNotFound = errors.New("id not found")
)

func (u *deleteUsecase) Delete(accountId string) error {
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
