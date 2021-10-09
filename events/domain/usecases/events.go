package usecases

import "events/domain/entities"

type CreateEvent interface {
	Create(event *entities.Event) (*entities.Event, error)
}

type GetEvents interface {
	GetAccountEvents(accountId string) ([]*entities.Event, error)
}
