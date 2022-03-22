package application

import (
	"events/domain/entities"
	"events/framework/db/repositories"
)

type CreateEvent struct {
	repository repositories.EventsRepository
}

// NewCreateEvent returns a pointer to a instace of a `CreateEvent` structure
func NewCreateEvent(repo repositories.EventsRepository) *CreateEvent {
	return &CreateEvent{
		repository: repo,
	}
}

func (c *CreateEvent) Create(params *CreateEventParams) (*entities.Event, error) {

	event, err := entities.NewEvent(
		params.Name, params.Description, params.OrganizerAccountId, params.AgeGroup,
		params.MaximumCapacity, nil, params.Duration, nil, params.Date)
	if err != nil {
		return nil, err
	}

	location := params.Location
	eventLocation, err := entities.NewEventLocation(
		location.Street, event.Id, location.District, location.State, location.City,
		location.PostalCode, location.Description, location.Number,
		location.Latitude, params.Location.Longitude)
	if err != nil {
		return nil, err
	}

	event.Location = *eventLocation

	c.setEventTyckets(event, params.TycketOptions)

	if err = c.repository.Create(event); err != nil {
		return nil, err
	}

	return event, nil
}

// setEventTyckets - creates an entity for all `TyketOption` and it's `TycketLot` and set those
// etities into the event
func (*CreateEvent) setEventTyckets(event *entities.Event, params []TycketOptionsParams) {
	var tycketOptions []entities.TycketOption

	for _, option := range params {

		tycketOption := entities.NewTycketOption(event.Id, option.Title, nil)

		lots := option.Lots
		var tycketLots []entities.TycketLot

		for _, lot := range lots {
			tycketLot := entities.NewTycketLot(lot.Number, tycketOption.Id, lot.TycketPrice, lot.TycketAmount)
			tycketLots = append(tycketLots, *tycketLot)
		}

		tycketOption.Lots = tycketLots
		tycketOptions = append(tycketOptions, *tycketOption)
	}

	event.TycketOptions = tycketOptions
}
