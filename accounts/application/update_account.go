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
	repository      db.Repository
	cacheRepository cache.Repository
}

func NewUpdateAccount(repository db.Repository, cacheRepo cache.Repository) *UpdateAccount {
	return &UpdateAccount{
		repository:      repository,
		cacheRepository: cacheRepo,
	}
}

func (c *UpdateAccount) Update(accountId string, data *UpdateAccountDTO) (*entities.Account, error) {
	account, err := c.repository.FindById(accountId)
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

	err = c.repository.Update(account)
	if err != nil {
		return nil, err
	}

	value, err := json.Marshal(account)
	if err != nil {
		return nil, err
	}

	err = c.cacheRepository.Set(account.Id, string(value), time.Minute*10)
	if err != nil {
		return nil, err
	}

	return account, nil
}
