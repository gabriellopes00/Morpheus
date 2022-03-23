package entities

import "events/domain"

type TicketLot struct {
	domain.Entity
	Number   int     `json:"number,omitempty"`
	Price    float64 `json:"price,omitempty"`
	Quantity int     `json:"quantity,omitempty"`
}
