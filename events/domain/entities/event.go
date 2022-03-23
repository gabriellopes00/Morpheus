package entities

import (
	"events/domain"
	domain_errors "events/domain/errors"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

type EventStatus string

const (
	StatusAvailable EventStatus = "available"
	StatusSoldOut   EventStatus = "sold_out"
	StatusFinished  EventStatus = "finished"
	StatusCanceled  EventStatus = "canceled"
)

type EventVisibility string

const (
	VisibilityPrivate     EventVisibility = "private"
	VisibilityPublic      EventVisibility = "private"
	VisibilityInvitedOnly EventVisibility = "invitedonly"
)

type Event struct {
	domain.Entity
	Name               string          `json:"name,omitempty"`
	Description        string          `json:"description,omitempty"`
	CoverUrl           string          `json:"cover_url,omitempty"`
	OrganizerAccountId string          `json:"organizer_account_id,omitempty"`
	AgeGroup           int             `json:"age_group,omitempty"`
	Status             EventStatus     `json:"status,omitempty"`
	Location           EventLocation   `json:"location" gorm:"foreignKey:EventId;references:Id"`
	Tickets            []Ticket        `json:"ticket_options,omitempty" gorm:"foreignKey:EventId;references:Id"`
	StartDateTime      time.Time       `json:"start_datetime,omitempty"`
	EndDateTime        time.Time       `json:"end_datetime,omitempty"`
	CategoryId         string          `json:"category,omitempty"`
	SubjectId          string          `json:"subject,omitempty"`
	Visibility         EventVisibility `json:"visibility,omitempty"`
}

func NewEvent(
	name, description string, organizerAccountId string,
	ageGroup, maximumCapacity int, location *EventLocation,
	duration int, tycketOptions []TycketOption, date string,
) (*Event, domain_errors.DomainErr) {
	var err error

	event := new(Event)
	event.Id = uuid.NewV4().String()
	event.Name = strings.TrimSpace(name)
	event.Description = strings.TrimSpace(description)

	event.StartDateTime, err = time.Parse(time.RFC3339, date)
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

	return nil
}
