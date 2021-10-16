package application

import (
	"events/domain/entities"
	"events/framework/db/repositories"
)

type FindEvents struct {
	Repository repositories.EventsRepository
}

func NewFindEvents(repo repositories.EventsRepository) *FindEvents {
	return &FindEvents{
		Repository: repo,
	}
}

func (u *FindEvents) FindAccountEvents(accountId string) ([]*entities.Event, error) {
	return u.Repository.FindAccountEvents(accountId)
}

func (u *FindEvents) FindEventById(eventId string) (*entities.Event, error) {
	return u.Repository.FindById(eventId)
}
