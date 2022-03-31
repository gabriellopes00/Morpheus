package entities

import (
	"events/domain"
	"time"

	uuid "github.com/satori/go.uuid"
)

type TicketOptionLot struct {
	domain.Entity
	Number   int     `json:"number,omitempty"`
	Price    float64 `json:"price,omitempty"`
	Quantity int     `json:"quantity,omitempty"`
}

func NewTicketOptionLot(number, quantity int, price float64) *TicketOptionLot {
	ticketOptionLot := new(TicketOptionLot)

	ticketOptionLot.Id = uuid.NewV4().String()
	ticketOptionLot.CreatedAt = time.Now()
	ticketOptionLot.UpdatedAt = time.Now()

	ticketOptionLot.Number = number
	ticketOptionLot.Price = price
	ticketOptionLot.Quantity = quantity

	return ticketOptionLot
}
