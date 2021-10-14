package db

import "tickets/domain/entities"

type TicketsRepository interface {
	Create(ticket *entities.Ticket) error
}
