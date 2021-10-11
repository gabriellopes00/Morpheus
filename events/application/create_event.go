package application

import (
	"events/domain/entities"
	"events/framework/db/repositories"
)

type createEventUsecase struct {
	Repository repositories.EventsRepository
}

func NewCreateEventUsecase(repo repositories.EventsRepository) *createEventUsecase {
	return &createEventUsecase{
		Repository: repo,
	}
}

type CreateEventParams struct {
	Name               string `json:"name,omitempty"`
	Description        string `json:"description,omitempty"`
	IsAvailable        bool   `json:"is_available,omitempty"`
	OrganizerAccountId string `json:"organizer_account_id,omitempty"`
	AgeGroup           int    `json:"age_group,omitempty"`
	MaximumCapacity    int    `json:"maximum_capacity,omitempty"`
	Status             string `json:"status,omitempty"`
	Location           struct {
		Street      string  `json:"street,omitempty"`
		District    string  `json:"district,omitempty"`
		State       string  `json:"state,omitempty"`
		City        string  `json:"city,omitempty"`
		Number      int     `json:"number,omitempty"`
		PostalCode  string  `json:"postal_code,omitempty"`
		Description string  `json:"description,omitempty"`
		Latitude    float64 `json:"latitude,omitempty"`
		Longitude   float64 `json:"longitude,omitempty"`
	} `json:"location,omitempty"`
	TicketPrice float32  `json:"ticket_price,omitempty"`
	Dates       []string `json:"dates,omitempty"`
}

func (c *createEventUsecase) Create(params *CreateEventParams) (*entities.Event, error) {

	location := params.Location
	eventLocation, err := entities.NewEventLocation(
		location.Street, location.District, location.State, location.City,
		location.PostalCode, location.Description, location.Number,
		location.Latitude, params.Location.Longitude)
	if err != nil {
		return nil, err
	}

	event, err := entities.NewEvent(
		params.Name, params.Description, params.IsAvailable, params.OrganizerAccountId, params.AgeGroup,
		params.MaximumCapacity, params.Status, eventLocation, params.TicketPrice, params.Dates)
	if err != nil {
		return nil, err
	}

	if err = c.Repository.Create(event); err != nil {
		return nil, err
	}

	return event, nil
}
