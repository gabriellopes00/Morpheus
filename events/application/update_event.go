package application

import (
	"errors"
	"events/domain/entities"
	"events/framework/db/repositories"
	"time"
)

type UpdateEvent struct {
	Repository repositories.EventsRepository
}

func NewUpdateEvent(repo repositories.EventsRepository) *UpdateEvent {
	return &UpdateEvent{
		Repository: repo,
	}
}

func (u *UpdateEvent) SetStatus(eventId string, status entities.EventStatus) error {
	return u.Repository.SetStatus(eventId, status)
}

type UpdateEventDTO struct {
	Name            string                 `json:"name,omitempty"`
	Description     string                 `json:"description,omitempty"`
	AgeGroup        int                    `json:"age_group,omitempty"`
	MaximumCapacity int                    `json:"maximum_capacity,omitempty"`
	Location        entities.EventLocation `json:"location"`
	Duration        int                    `json:"duration,omitempty"`
	Date            time.Time              `json:"date,omitempty"`
}

func (u *UpdateEvent) UpdateData(eventId string, data *UpdateEventDTO) (*entities.Event, error) {
	event, err := u.Repository.FindById(eventId)
	if err != nil {
		return nil, err
	}

	if event == nil {
		return nil, errors.New("event with given id does not exists")
	}

	event.Name = data.Name
	event.Description = data.Description
	event.AgeGroup = data.AgeGroup
	event.MaximumCapacity = data.MaximumCapacity
	event.Location = data.Location
	event.Duration = data.Duration

	if time.Until(event.Date) <= time.Hour*24 {
		return nil, errors.New("the event date cannot be updated less than 24 hours before the event starts")
	}

	event.Date = data.Date

	err = u.Repository.Update(event)
	if err != nil {
		return nil, err
	}

	return event, nil
}
