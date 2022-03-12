package application

import (
	"accounts/pkg/cache"
	"accounts/pkg/db"
	"errors"
)

type DeleteAccount struct {
	Repository      db.Repository
	CacheRepository cache.Repository
}

func NewDeleteAccount(Repository db.Repository, CacheRepo cache.Repository) *DeleteAccount {
	return &DeleteAccount{
		Repository:      Repository,
		CacheRepository: CacheRepo,
	}
}

var (
	ErrIdNotFound = errors.New("id not found")
)

func (d *DeleteAccount) Delete(accountId string) error {
	existingAccount, err := d.Repository.ExistsId(accountId)
	if err != nil {
		return err
	}

	if !existingAccount {
		return ErrIdNotFound
	}

	err = d.Repository.Delete(accountId)
	if err != nil {
		return err
	}

	err = d.CacheRepository.Delete(accountId)
	if err != nil {
		return err
	}

	return err
}
