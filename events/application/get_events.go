package application

import (
	"events/domain/entities"
	"events/framework/db/repositories"
)

type GetEvents struct {
	Repository repositories.EventsRepository
}

func NewGetEvents(repo repositories.EventsRepository) *GetEvents {
	return &GetEvents{
		Repository: repo,
	}
}

func (u *GetEvents) GetAccountEvents(accountId string) ([]*entities.Event, error) {

	return u.Repository.GetAccountEvents(accountId)
}
