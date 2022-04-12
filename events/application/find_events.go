package application

import (
	"errors"
	"events/domain/entities"
	"events/framework/db/repositories"
)

type FindEvents struct {
	repository repositories.EventsRepository
}

func NewFindEvents(repo repositories.EventsRepository) *FindEvents {
	return &FindEvents{
		repository: repo,
	}
}

func (u *FindEvents) FindAccountEvents(accountId string) ([]entities.Event, error) {
	return u.repository.FindAccountEvents(accountId)
}

func (u *FindEvents) FindEventById(eventId string) (*entities.Event, error) {
	return u.repository.FindById(eventId)
}

func (u *FindEvents) FindAll(state string, month, ageGroup, limit, offset int) ([]entities.Event, error) {
	if len(state) != 2 {
		return nil, errors.New("invalid state abbreviation")
	}

	if month < 0 || month > 12 {
		return nil, errors.New("invalid month")
	}

	if limit < 1 && limit > 30 {
		return nil, errors.New("invalid results limit")
	}

	switch ageGroup {
	case 0, 10, 12, 14, 16, 18:
		break
	default:
		return nil, errors.New("age group must be: 0, 10, 12, 14, 16 or 18")

	}

	return u.repository.FindAll(state, month, ageGroup, limit, offset)
}
