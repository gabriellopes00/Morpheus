package application

import (
	"errors"
	"tickets/domain/entities"
	"tickets/framework/db"
)

type ticketGenerator struct {
	ticketsRepo db.TicketsRepository
	eventsRepo  db.EventsRepository
}

func NewTicketGenerator(
	ticketsRepository db.TicketsRepository,
	eventsRepository db.EventsRepository) *ticketGenerator {
	return &ticketGenerator{
		ticketsRepo: ticketsRepository,
		eventsRepo:  eventsRepository,
	}
}

func (t *ticketGenerator) GenerateTicket(eventId, ownerAccountId string) (*entities.Ticket, error) {
	event, err := t.eventsRepo.FindById(eventId)
	if err != nil {
		return nil, err
	}

	if event == nil {
		return nil, errors.New("event not found")
	}

	ticket := entities.NewTicket(eventId, ownerAccountId, event.Date)

	if err = t.ticketsRepo.Create(ticket); err != nil {
		return nil, err
	}

	return ticket, nil
}
