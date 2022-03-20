package entities

import (
	domain_errors "events/domain/errors"
	"strings"
	"time"

	gouuid "github.com/satori/go.uuid"
)

type EventStatus string

const (
	StatusAvailable EventStatus = "available"
	StatusSoldOut   EventStatus = "sold_out"
	StatusFinished  EventStatus = "finished"
	StatusCanceled  EventStatus = "canceled"
)

type Event struct {
	Id                 string         `json:"id,omitempty" gorm:"primaryKey"`
	Name               string         `json:"name,omitempty"`
	Description        string         `json:"description,omitempty"`
	OrganizerAccountId string         `json:"organizer_account_id,omitempty"`
	AgeGroup           int            `json:"age_group,omitempty"`
	MaximumCapacity    int            `json:"maximum_capacity,omitempty"`
	Status             EventStatus    `json:"status,omitempty"`
	Location           EventLocation  `json:"location" gorm:"foreignKey:Id"`
	Duration           int            `json:"duration,omitempty"`
	TycketOptions      []TycketOption `json:"tycket_options,omitempty" gorm:"foreignKey:Id"`
	Date               time.Time      `json:"date,omitempty"`
	CreatedAt          time.Time      `json:"created_at,omitempty"`
	UpdatedAt          time.Time      `json:"updated_at,omitempty"`
}

func NewEvent(
	name, description string, organizerAccountId string,
	ageGroup, maximumCapacity int, location *EventLocation,
	duration int, tycketOptions []TycketOption, date string,
) (*Event, domain_errors.DomainErr) {
	var err error

	event := &Event{
		Id:                 gouuid.NewV4().String(),
		Name:               strings.TrimSpace(name),
		Description:        strings.TrimSpace(description),
		OrganizerAccountId: organizerAccountId,
		AgeGroup:           ageGroup,
		MaximumCapacity:    maximumCapacity,
		Location:           EventLocation{},
		Status:             StatusAvailable,
		TycketOptions:      tycketOptions, // TODO: validate maximum capacity
		Duration:           duration,
		UpdatedAt:          time.Now().Local(),
		CreatedAt:          time.Now().Local(),
	}

	event.Date, err = time.Parse(time.RFC3339, date)
	if err != nil {
		return nil, domain_errors.NewValidationError(
			`Event's dates must be in "RFC3339" format`,
			"Date", date)
	}

	if err = event.validate(); err != nil {
		return nil, err
	}

	return event, nil
}

func (e *Event) validate() domain_errors.DomainErr {
	if len(e.Name) == 4 || len(e.Name) > 255 {
		return domain_errors.NewValidationError(
			"Events name must have at least of 4 characters and at most of 255",
			"name",
			e.Name)
	}

	switch e.AgeGroup {
	case 0, 10, 12, 14, 16, 18:
		break
	default:
		{
			return domain_errors.NewValidationError(
				"Events' age group must be: 0, 10, 12, 14, 16 or 18",
				"AgeGroup",
				e.AgeGroup)
		}
	}

	if e.MaximumCapacity <= 0 {
		return domain_errors.NewValidationError(
			"Events' maximum capacity must be greather than 0",
			"MaximumAge",
			e.MaximumCapacity)
	}

	if e.Duration <= 0 {
		return domain_errors.NewValidationError(
			"Events' duration must be greather than 0 minutes",
			"Duration",
			e.Duration)
	}

	return nil
}
