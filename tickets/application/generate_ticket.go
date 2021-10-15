package application

import (
	"errors"
	"tickets/domain/entities"
	"tickets/framework/db/repositories"
)

type ticketGenerator struct {
	ticketsRepo repositories.TicketsRepository
	eventsRepo  repositories.EventsRepository
}

func NewTicketGenerator(
	ticketsRepository repositories.TicketsRepository,
	eventsRepository repositories.EventsRepository) *ticketGenerator {
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
