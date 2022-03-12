package application

import (
	"accounts/domain/entities"
	"accounts/pkg/cache"
	"accounts/pkg/db"
	"encoding/json"
	"time"
)

type UpdateAccountDTO struct {
	Name      string `json:"name,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	BirthDate string `json:"birth_date,omitempty"`
}

type UpdateAccount struct {
	Repository      db.Repository
	CacheRepository cache.Repository
}

func NewUpdateAccount(Repository db.Repository, CacheRepo cache.Repository) *UpdateAccount {
	return &UpdateAccount{
		Repository:      Repository,
		CacheRepository: CacheRepo,
	}
}

func (c *UpdateAccount) Update(accountId string, data *UpdateAccountDTO) (*entities.Account, error) {
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

	value, err := json.Marshal(account)
	if err != nil {
		return nil, err
	}

	err = c.CacheRepository.Set(account.Id, string(value), time.Minute*10)
	if err != nil {
		return nil, err
	}

	return account, nil
}
