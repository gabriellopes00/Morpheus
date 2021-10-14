package db

import "tickets/domain/entities"

type EventsRepository interface {
	FindById(eventId string) (*entities.Event, error)
	Create(event *entities.Event) error
}
