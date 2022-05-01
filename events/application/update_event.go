package application

import (
	"encoding/json"
	"errors"
	"events/domain/entities"
	"events/framework/db/repositories"
	"time"

	"github.com/imdario/mergo"
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
	Name          string         `json:"name,omitempty"`
	Description   string         `json:"description,omitempty"`
	CoverUrl      string         `json:"cover_url,omitempty"`
	AgeGroup      int            `json:"age_group,omitempty"`
	Location      LocationParams `json:"location,omitempty"`
	StartDateTime string         `json:"start_datetime,omitempty" gorm:"column:start_datetime"`
	EndDateTime   string         `json:"end_datetime,omitempty" gorm:"column:end_datetime"`
	CategoryId    string         `json:"category,omitempty" gorm:"foreignKey:CategoryId"`
	SubjectId     string         `json:"subject,omitempty" gorm:"foreignKey:SubjectId"`
	Visibility    string         `json:"visibility,omitempty"`
	Status        string         `json:"status,omitempty"`
}

func (u *UpdateEvent) UpdateData(eventId string, data *UpdateEventDTO) (*entities.Event, error) {
	event, err := u.Repository.FindById(eventId)
	if err != nil {
		return nil, err
	}

	if event == nil {
		return nil, errors.New("event with given id does not exists")
	}

	j, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	newEvent := entities.Event{}
	err = json.Unmarshal(j, &newEvent)
	if err != nil {
		return event, err
	}

	if err := mergo.Merge(event, newEvent, mergo.WithOverride); err != nil {
		return event, err
	}

	event.UpdatedAt = time.Now()

	// if time.Duration(time.Until(event.StartDateTime).Hours()) <= time.Hour*24 {
	// 	return nil, errors.New("the event date cannot be updated less than 24 hours before the event starts")
	// }

	err = u.Repository.Update(event)
	if err != nil {
		return nil, err
	}

	return event, nil
}
