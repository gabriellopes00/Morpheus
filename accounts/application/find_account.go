package application

import (
	"accounts/domain/entities"
	"accounts/pkg/cache"
	"accounts/pkg/db"
	"encoding/json"
	"time"
)

type FindAccount struct {
	repository      db.Repository
	cacheRepository cache.Repository
}

func NewFindAccount(repo db.Repository, cacheRepo cache.Repository) *FindAccount {
	return &FindAccount{
		repository:      repo,
		cacheRepository: cacheRepo,
	}
}

func (f *FindAccount) FindById(accountId string) (*entities.Account, error) {

	data, err := f.cacheRepository.Get(accountId)
	if err != nil {
		return nil, err
	}

	if data == "" {
		result, err := f.repository.FindById(accountId)
		if err != nil {
			return nil, err
		}

		if result == nil {
			return nil, nil
		}

		value, _ := json.Marshal(result)

		err = f.cacheRepository.Set(result.Id, string(value), time.Minute*10)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	var account entities.Account
	json.Unmarshal([]byte(data), &account)

	return &account, nil
}
