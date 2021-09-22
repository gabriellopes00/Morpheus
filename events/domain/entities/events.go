package entities

import (
	"events/domain/errors"
	"strings"
	"time"

	gouuid "github.com/satori/go.uuid"
)

type Event struct {
	Id                 string    `json:"id,omitempty"`
	Name               string    `json:"name,omitempty"`
	Description        string    `json:"description,omitempty"`
	IsAvailable        bool      `json:"is_available,omitempty"`
	OrganizerAccountId string    `json:"organizer_account_id,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
}

func NewEvent(name, description, organizerAccountId string, isAvailable bool) (*Event, errors.DomainErr) {
	event := &Event{
		Id:                 gouuid.NewV4().String(),
		Name:               strings.TrimSpace(name),
		Description:        strings.TrimSpace(description),
		IsAvailable:        isAvailable,
		OrganizerAccountId: organizerAccountId,
		CreatedAt:          time.Now().Local(),
	}

	err := validate(event)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func validate(e *Event) errors.DomainErr {
	if len(e.Name) == 4 || len(e.Name) > 255 {
		return errors.NewValidationError(
			"Events name must have at least of 4 characters and at most of 255",
			"name")
	}

	if len(e.Description) == 4 {
		return errors.NewValidationError(
			"Events descriptions must have at least of 4 characters",
			"description")
	}

	return nil
}
