package application

import (
	"errors"
	"events/domain/entities"
	"events/framework/db/repositories"
	"fmt"
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
}

func (u *UpdateEvent) UpdateData(eventId string, data *UpdateEventDTO) (*entities.Event, error) {
	event, err := u.Repository.FindById(eventId)
	if err != nil {
		return nil, err
	}

	if event == nil {
		return nil, errors.New("event with given id does not exists")
	}

	err = mergo.Merge(event, data, mergo.WithOverride) // must be the same type
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println()
	fmt.Println(event)
	fmt.Println()

	// event.Name = data.Name || ""
	// event.Description = data.Description
	// event.AgeGroup = data.AgeGroup
	// event.CategoryId = data.CategoryId
	// event.SubjectId = data.SubjectId
	// event.Visibility = entities.EventVisibility(data.Visibility)

	if time.Until(event.StartDateTime) <= time.Hour*24 {
		return nil, errors.New("the event date cannot be updated less than 24 hours before the event starts")
	}

	// event.StartDateTime = data.StartDateTime
	// event.EndDateTime = data.EndDateTime
	// event.Location.Street = data.Location.Street
	// event.Location.District = data.Location.District
	// event.Location.State = data.Location.State
	// event.Location.City = data.Location.City
	// event.Location.Number = data.Location.Number
	// event.Location.PostalCode = data.Location.PostalCode
	// event.Location.Description = data.Location.Description
	// event.Location.Latitude = data.Location.Latitude
	// event.Location.Longitude = data.Location.Longitude

	err = u.Repository.Update(event)
	if err != nil {
		return nil, err
	}

	return event, nil
}
