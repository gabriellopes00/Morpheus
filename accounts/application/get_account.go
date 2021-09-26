package usecases

import (
	"accounts/domain"
	"accounts/interfaces"
)

type getAccount struct {
	Repository interfaces.Repository
}

func NewGetAccount(repo interfaces.Repository) *getAccount {
	return &getAccount{
		Repository: repo,
	}
}

func (g *getAccount) GetById(accountId string) (*domain.Account, error) {
	return g.Repository.FindById(accountId)
}
