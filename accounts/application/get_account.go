package usecases

import (
	"accounts/domain"
	"accounts/pkg/db"
)

type getAccount struct {
	Repository db.Repository
}

func NewGetAccount(repo db.Repository) *getAccount {
	return &getAccount{
		Repository: repo,
	}
}

func (g *getAccount) GetById(accountId string) (*domain.Account, error) {
	return g.Repository.FindById(accountId)
}
