package application

import (
	"tickets/domain/entities"
	"tickets/framework/db"
	"time"
)

type createEvent struct {
	repo db.EventsRepository
}

func NewCreateEvent(repository db.EventsRepository) *createEvent {
	return &createEvent{
		repo: repository,
	}
}

func (c *createEvent) Create(eventId string, date time.Time) (*entities.Event, error) {
	event := entities.NewEvent(eventId, date)
	if err := c.repo.Create(event); err != nil {
		return nil, err
	}

	return event, nil

}
