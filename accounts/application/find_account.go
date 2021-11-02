package application

import (
	"accounts/domain/entities"
	"accounts/pkg/db"
)

type findAccount struct {
	Repository db.Repository
}

func NewFindAccount(repo db.Repository) *findAccount {
	return &findAccount{
		Repository: repo,
	}
}

func (g *findAccount) FindById(accountId string) (*entities.Account, error) {
	return g.Repository.FindById(accountId)
}
