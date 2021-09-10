package entities

import (
	"events/domain/errors"
	"math"
	"strings"
	"time"

	gouuid "github.com/satori/go.uuid"
)

type Event struct {
	Id          string    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Price       float64   `json:"price,omitempty"`
	IsAvailable bool      `json:"is_available,omitempty"`
	OwnerId     string    `json:"owner_id,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewEvent(title, description, ownerId string, price float64, isAvailable bool) (*Event, errors.DomainErr) {
	event := &Event{
		Id:          gouuid.NewV4().String(),
		Title:       strings.TrimSpace(title),
		Description: strings.TrimSpace(description),
		Price:       math.Round(price*100) / 100,
		IsAvailable: isAvailable,
		OwnerId:     ownerId,
		CreatedAt:   time.Now().Local(),
	}

	err := validate(event)
	if err != nil {
		return nil, err
	}

	return event, nil
}

func validate(e *Event) errors.DomainErr {
	if len(e.Title) == 4 || len(e.Title) > 255 {
		return errors.NewValidationError(
			"Events title must have at least of 4 characters and at most of 255",
			"title")
	}

	if len(e.Description) == 4 {
		return errors.NewValidationError(
			"Events descriptions must have at least of 4 characters",
			"description")
	}

	if e.Price < 0 {
		return errors.NewValidationError(
			"Events price cannot be less than 0",
			"price")
	}

	return nil
}
