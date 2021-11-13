package application

import (
	"errors"
	"events/domain/entities"
	"events/framework/db/repositories"
	"time"
)

type updateEventUsecase struct {
	Repository repositories.EventsRepository
}

func NewUpdateEventUsecase(repo repositories.EventsRepository) *updateEventUsecase {
	return &updateEventUsecase{
		Repository: repo,
	}
}

func (u *updateEventUsecase) UpdateStatus(eventId string, status interface{}) error {
	stts, ok := status.(entities.EventStatus)
	if !ok {
		return errors.New("invalid status")
	}

	eventExists, err := u.Repository.ExistsId(eventId)
	if err != nil {
		return err
	}

	if !eventExists {
		return errors.New("event with given id does not exists")
	}

	return u.Repository.SetStatus(eventId, stts)
}

type UpdateEventDTO struct {
	Name            string                 `json:"name,omitempty"`
	Description     string                 `json:"description,omitempty"`
	IsAvailable     bool                   `json:"is_available,omitempty"`
	AgeGroup        int                    `json:"age_group,omitempty"`
	MaximumCapacity int                    `json:"maximum_capacity,omitempty"`
	Location        entities.EventLocation `json:"location"`
	Duration        int                    `json:"duration,omitempty"`
	TicketPrice     float32                `json:"ticket_price,omitempty"`
	Date            time.Time              `json:"date,omitempty"`
}

func (u *updateEventUsecase) UpdateData(eventId string, data *UpdateEventDTO) (*entities.Event, error) {
	event, err := u.Repository.FindById(eventId)
	if err != nil {
		return nil, err
	}

	if event == nil {
		return nil, errors.New("event with given id does not exists")
	}

	event.Name = data.Name
	event.Description = data.Description
	event.IsAvailable = data.IsAvailable
	event.AgeGroup = data.AgeGroup
	event.MaximumCapacity = data.MaximumCapacity
	event.Location = data.Location
	event.Duration = data.Duration
	event.TicketPrice = data.TicketPrice
	event.Date = data.Date

	err = u.Repository.Update(event)
	if err != nil {
		return nil, err
	}

	return event, nil
}
