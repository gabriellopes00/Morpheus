package application

import (
	"errors"
	"events/domain/entities"
	"events/framework/db/repositories"
	"events/framework/geocode"
)

type FindEvents struct {
	repository      repositories.EventsRepository
	geocodeProvider geocode.GeocodeProvider
}

func NewFindEvents(repo repositories.EventsRepository, geocodeProvider geocode.GeocodeProvider) *FindEvents {
	return &FindEvents{
		repository:      repo,
		geocodeProvider: geocodeProvider,
	}
}

func (u *FindEvents) FindAccountEvents(accountId string) ([]entities.Event, error) {
	return u.repository.FindAccountEvents(accountId, false)
}

func (u *FindEvents) FindEventById(eventId string, deepFind bool) (*entities.Event, error) {
	return u.repository.FindById(eventId, deepFind)
}

func (u *FindEvents) FindNearbyEvents(latitude, longitude string) ([]entities.Event, error) {

	data, err := u.geocodeProvider.Reverse(latitude, longitude)
	if err != nil {
		return nil, err
	}

	// fetch event by city in cache

	events, err := u.repository.FindByLocation("", data.City) // cache it
	if err != nil {
		return nil, err
	}

	return events, nil
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
