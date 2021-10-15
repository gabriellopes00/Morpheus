package application

import (
	"tickets/domain/entities"
	"tickets/framework/db/repositories"
	"time"
)

type createEvent struct {
	repo repositories.EventsRepository
}

func NewCreateEvent(repository repositories.EventsRepository) *createEvent {
	return &createEvent{
		repo: repository,
	}
}

func (c *createEvent) Create(eventId string, date time.Time) error {
	event := entities.NewEvent(eventId, date)
	if err := c.repo.Create(event); err != nil {
		return err
	}

	return nil

}
