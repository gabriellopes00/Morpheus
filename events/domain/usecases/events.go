package usecases

import (
	"events/application"
	"events/domain/entities"
)

type CreateEvent interface {
	Create(event *application.CreateEventParams) (*entities.Event, error)
}

type FindEvents interface {
	FindAccountEvents(accountId string) ([]*entities.Event, error)
	FindEventById(eventId string) (*entities.Event, error)
}
