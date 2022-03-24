package entities

import (
	"events/domain"
	domain_errors "events/domain/errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Ticket struct {
	domain.Entity
	Title               string      `json:"title,omitempty"`
	Description         string      `json:"description,omitempty"`
	SalesStartDateTime  time.Time   `json:"sales_start_datetime,omitempty"`
	SalesEndDateTime    time.Time   `json:"sales_end_datetime,omitempty"`
	MinimumBuysQuantity int         `json:"maximum_buys_quantity,omitempty"`
	MaximumBuysQuantity int         `json:"minimum_buys_quantity,omitempty"`
	Lots                []TicketLot `json:"lots,omitempty"`
}

func NewTicket(
	title, description, salesStartDateTime, salesEndDateTime string,
	maximumBuysQuantity, minimumBuysQuantity int, lots []TicketLot) (*Ticket, error) {

	var err error

	ticket := new(Ticket)

	ticket.Id = uuid.NewV4().String()
	ticket.CreatedAt = time.Now()
	ticket.UpdatedAt = time.Now()

	ticket.Title = title
	ticket.Description = description
	ticket.MaximumBuysQuantity = maximumBuysQuantity
	ticket.MinimumBuysQuantity = minimumBuysQuantity
	ticket.Lots = lots

	ticket.SalesStartDateTime, err = time.Parse(time.RFC3339, salesStartDateTime)
	if err != nil {
		return nil, domain_errors.NewValidationError(
			`Ticket's dates must be in "RFC3339" format`,
			"SalesStartDateTime", salesStartDateTime)
	}

	ticket.SalesEndDateTime, err = time.Parse(time.RFC3339, salesEndDateTime)
	if err != nil {
		return nil, domain_errors.NewValidationError(
			`Ticket's dates must be in "RFC3339" format`,
			"SalesEndDateTime", salesEndDateTime)
	}

	return ticket, nil
}
