package application

import (
	"accounts/domain/entities"
	"accounts/pkg/cache"
	"accounts/pkg/db"
	"encoding/json"
	"time"
)

type FindAccount struct {
	Repository      db.Repository
	CacheRepository cache.Repository
}

func NewFindAccount(repo db.Repository, cacheRepo cache.Repository) *FindAccount {
	return &FindAccount{
		Repository:      repo,
		CacheRepository: cacheRepo,
	}
}

func (f *FindAccount) FindById(accountId string) (*entities.Account, error) {

	data, err := f.CacheRepository.Get(accountId)
	if err != nil {
		return nil, err
	}

	if data == "" {
		result, err := f.Repository.FindById(accountId)
		if err != nil {
			return nil, err
		}

		if result == nil {
			return nil, nil
		}

		value, err := json.Marshal(result)
		if err != nil {
			return nil, err
		}

		err = f.CacheRepository.Set(result.Id, string(value), time.Minute*10)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	var account entities.Account
	err = json.Unmarshal([]byte(data), &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}
