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
		params.Name, params.Description, params.CoverUrl, params.OrganizerAccountId,
		params.AgeGroup, nil, nil, params.StartDateTime, params.EndDateTime, params.CategoryId,
		params.SubjectId, params.Visibility)
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

	event.Location = eventLocation

	c.setEventTickets(event, params.Tickets)

	if err = c.repository.Create(event); err != nil {
		return nil, err
	}

	return event, nil
}

// setEventTickets - creates an entity for all `Ticket` and it's `TicketLot` and set those
// entities into the event
func (*CreateEvent) setEventTickets(event *entities.Event, params []TicketParams) error {
	var tickets []entities.Ticket

	for _, param := range params {

		lots := param.Lots
		var ticketLots []entities.TicketLot

		for _, lot := range lots {
			ticketLot := entities.NewTicketLot(lot.Number, lot.Quantity, lot.Price)
			ticketLots = append(ticketLots, *ticketLot)
		}

		ticket, err := entities.NewTicket(param.Title, param.Description, param.SalesStartDateTime,
			param.SalesEndDateTime, param.MaximumBuysQuantity, param.MinimumBuysQuantity, ticketLots)
		if err != nil {
			return err
		}

		tickets = append(tickets, *ticket)
	}

	event.Tickets = tickets
	return nil
}
