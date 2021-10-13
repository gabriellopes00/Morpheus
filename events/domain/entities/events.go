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
	Id                 string        `json:"id,omitempty"`
	Name               string        `json:"name,omitempty"`
	Description        string        `json:"description,omitempty"`
	IsAvailable        bool          `json:"is_available,omitempty"`
	OrganizerAccountId string        `json:"organizer_account_id,omitempty"`
	AgeGroup           int           `json:"age_group,omitempty"`
	MaximumCapacity    int           `json:"maximum_capacity,omitempty"`
	Status             EventStatus   `json:"status,omitempty"`
	Location           EventLocation `json:"location"`
	Duration           int           `json:"duration,omitempty"`
	TicketPrice        float32       `json:"ticket_price,omitempty"`
	Date               time.Time     `json:"date,omitempty"`
	UpdatedAt          time.Time     `json:"updated_at,omitempty"`
	CreatedAt          time.Time     `json:"created_at,omitempty"`
	DeletedAt          time.Time     `json:"deleted_at,omitempty"`
}

func NewEvent(
	name, description string, isAvailable bool, organizerAccountId string,
	ageGroup, maximumCapacity int, status string, location *EventLocation,
	duration int, ticketPrice float32, date string,
) (*Event, domain_errors.DomainErr) {
	var err error

	event := &Event{
		Id:                 gouuid.NewV4().String(),
		Name:               strings.TrimSpace(name),
		Description:        strings.TrimSpace(description),
		IsAvailable:        isAvailable,
		OrganizerAccountId: organizerAccountId,
		TicketPrice:        ticketPrice,
		AgeGroup:           ageGroup,
		MaximumCapacity:    maximumCapacity,
		Location:           *location,
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

	switch status {
	case string(StatusAvailable):
		event.Status = StatusAvailable
	case string(StatusCanceled):
		event.Status = StatusCanceled
	case string(StatusFinished):
		event.Status = StatusFinished
	case string(StatusSoldOut):
		event.Status = StatusSoldOut
	default:
		return nil, domain_errors.NewValidationError(
			`Events' status must be "available", "sold_out", "canceled" or "finished"`,
			"Status", status)
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

	if e.TicketPrice <= 0 {
		return domain_errors.NewValidationError(
			"Events' ticket price must be greather than 0",
			"TicketPrice",
			e.TicketPrice)
	}

	if e.Duration <= 0 {
		return domain_errors.NewValidationError(
			"Events' duration must be greather than 0 minutes",
			"Duration",
			e.Duration)
	}

	return nil
}
