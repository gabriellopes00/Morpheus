package application

import (
	"accounts/domain/entities"
	"accounts/domain/usecases"
	"accounts/pkg/db"
	"time"
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
	account, err := c.Repository.FindById(accountId)
	if err != nil {
		return nil, err
	}

	if account == nil {
		return nil, ErrIdNotFound
	}

	account.Name = data.Name
	account.AvatarUrl = data.AvatarUrl
	account.BirthDate, _ = time.Parse(time.RFC3339, data.BirthDate)
	account.UpdatedAt = time.Now()

	err = c.Repository.Update(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}
