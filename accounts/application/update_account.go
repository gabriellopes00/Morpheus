package application

import (
	"accounts/domain/entities"
	"accounts/domain/usecases"
	"accounts/pkg/db"
)

type updateAccount struct {
	Repository db.Repository
}

func NewUpdateAccount(Repository db.Repository) *updateAccount {
	return &updateAccount{
		Repository: Repository,
	}
}

func (c *updateAccount) Update(accountId string, data *usecases.UpdateAccountDTO) (*entities.Account, error) {
	existingAccount, err := c.Repository.ExistsId(accountId)
	if err != nil {
		return nil, err
	}

	if !existingAccount {
		return nil, ErrIdNotFound
	}

	account, err := c.Repository.Update(accountId, data)
	if err != nil {
		return nil, err
	}

	return account, nil
}
