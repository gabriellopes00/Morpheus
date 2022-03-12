package application

import (
	"accounts/domain/entities"
	"accounts/pkg/db"
)

type FindAccount struct {
	Repository db.Repository
}

func NewFindAccount(repo db.Repository) *FindAccount {
	return &FindAccount{
		Repository: repo,
	}
}

func (g *FindAccount) FindById(accountId string) (*entities.Account, error) {
	return g.Repository.FindById(accountId)
}
