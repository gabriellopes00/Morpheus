package usecases

import (
	"events/application"
	"events/domain/entities"
)

type CreateEvent interface {
	Create(event *application.CreateEventParams) (*entities.Event, error)
}

type GetEvents interface {
	GetAccountEvents(accountId string) ([]*entities.Event, error)
}
