package entities

import (
	"events/domain"
	domain_errors "events/domain/errors"
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
	VisibilityPublic      EventVisibility = "public"
	VisibilityInvitedOnly EventVisibility = "invited_only"
)

type Event struct {
	domain.Entity
	Name               string          `json:"name,omitempty"`
	Description        string          `json:"description,omitempty"`
	CoverUrl           string          `json:"cover_url,omitempty"`
	OrganizerAccountId string          `json:"organizer_account_id,omitempty"`
	AgeGroup           int             `json:"age_group,omitempty"`
	Status             EventStatus     `json:"status,omitempty"`
	Location           *EventLocation  `json:"location" gorm:"foreignKey:EventId;references:Id"`
	Tickets            []Ticket        `json:"ticket_options,omitempty" gorm:"foreignKey:EventId;references:Id"`
	StartDateTime      time.Time       `json:"start_datetime,omitempty"`
	EndDateTime        time.Time       `json:"end_datetime,omitempty"`
	CategoryId         string          `json:"category,omitempty" gorm:"foreignKey:CategoryId"`
	SubjectId          string          `json:"subject,omitempty" gorm:"foreignKey:SubjectId"`
	Visibility         EventVisibility `json:"visibility,omitempty"`
}

func NewEvent(
	name, description, coverUrl, organizerAccountId string,
	ageGroup int, location *EventLocation,
	tickets []Ticket, startDateTime, endDateTime string, categoryId,
	subjectId, visibility string,
) (*Event, domain_errors.DomainErr) {
	var err error

	event := new(Event)

	event.Id = uuid.NewV4().String()
	event.CreatedAt = time.Now()
	event.UpdatedAt = time.Now()

	event.Name = name
	event.Description = description
	event.CoverUrl = coverUrl
	event.OrganizerAccountId = organizerAccountId
	event.AgeGroup = ageGroup
	event.Location = location
	event.Tickets = tickets
	event.CategoryId = categoryId
	event.SubjectId = subjectId

	event.StartDateTime, err = time.Parse(time.RFC3339, startDateTime)
	if err != nil {
		return nil, domain_errors.NewValidationError(
			`Event's dates must be in "RFC3339" format`,
			"StartDateTime", startDateTime)
	}

	event.EndDateTime, err = time.Parse(time.RFC3339, endDateTime)
	if err != nil {
		return nil, domain_errors.NewValidationError(
			`Event's dates must be in "RFC3339" format`,
			"EndDateTime", endDateTime)
	}

	switch visibility {
	case string(VisibilityPrivate), string(VisibilityPublic), string(VisibilityInvitedOnly):
		event.Visibility = EventVisibility(visibility)
	default:
		{
			return nil, domain_errors.NewValidationError(
				"Events' visibility must be: private, public or invited_only",
				"Visibility", visibility)
		}
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

	return nil
}
