package entities

import (
	"events/domain"
	"time"

	uuid "github.com/satori/go.uuid"
)

type TicketLot struct {
	domain.Entity
	Number   int     `json:"number,omitempty"`
	Price    float64 `json:"price,omitempty"`
	Quantity int     `json:"quantity,omitempty"`
}

func NewTicketLot(number, quantity int, price float64) *TicketLot {
	ticketLot := new(TicketLot)

	ticketLot.Id = uuid.NewV4().String()
	ticketLot.CreatedAt = time.Now()
	ticketLot.UpdatedAt = time.Now()

	ticketLot.Number = number
	ticketLot.Price = price
	ticketLot.Quantity = quantity

	return ticketLot
}
