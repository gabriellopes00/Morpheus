package entities

import (
	"events/domain"
	"time"

	uuid "github.com/satori/go.uuid"
)

type TicketOptionLot struct {
	domain.Entity
	Number         int     `json:"number,omitempty"`
	TicketOptionId string  `json:"ticket_option_id,omitempty" gorm:"column:event_ticket_option_id"`
	Price          float64 `json:"price,omitempty"`
	Quantity       int     `json:"quantity,omitempty"`
}

func NewTicketOptionLot(number, quantity int, price float64, ticketOptionId string) *TicketOptionLot {
	ticketOptionLot := new(TicketOptionLot)

	ticketOptionLot.Id = uuid.NewV4().String()
	ticketOptionLot.CreatedAt = time.Now()
	ticketOptionLot.UpdatedAt = time.Now()

	ticketOptionLot.Number = number
	ticketOptionLot.TicketOptionId = ticketOptionId
	ticketOptionLot.Price = price
	ticketOptionLot.Quantity = quantity

	return ticketOptionLot
}

// gorm required
func (TicketOptionLot) TableName() string {
	return "event_ticket_options_lots"
}
