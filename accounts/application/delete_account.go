package application

import (
	"accounts/pkg/cache"
	"accounts/pkg/db"
	"errors"
)

type DeleteAccount struct {
	repository      db.Repository
	cacheRepository cache.Repository
}

func NewDeleteAccount(repository db.Repository, cacheRepo cache.Repository) *DeleteAccount {
	return &DeleteAccount{
		repository:      repository,
		cacheRepository: cacheRepo,
	}
}

var (
	ErrIdNotFound = errors.New("id not found")
)

func (d *DeleteAccount) Delete(accountId string) error {
	existingAccount, err := d.repository.ExistsId(accountId)
	if err != nil {
		return err
	}

	if !existingAccount {
		return ErrIdNotFound
	}

	err = d.repository.Delete(accountId)
	if err != nil {
		return err
	}

	err = d.cacheRepository.Delete(accountId)
	if err != nil {
		return err
	}

	return err
}
