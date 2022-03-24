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

	err = ticket.validate()
	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (ticket *Ticket) validate() error {
	if len(ticket.Title) < 1 || len(ticket.Title) > 255 {
		return domain_errors.NewValidationError(
			"Ticket's title must have at least of 4 characters and at most of 255",
			"title", ticket.Title)
	}

	if ticket.MinimumBuysQuantity < 1 {
		return domain_errors.NewValidationError(
			"Minimum buys tickets must be greather than 1",
			"MinimumBuysQuantity", ticket.MinimumBuysQuantity)
	}

	if ticket.MaximumBuysQuantity < ticket.MinimumBuysQuantity {
		return domain_errors.NewValidationError(
			"Maximum buys tickets must be greather than Minimum buys tickets",
			"MaximumBuysQuantity", ticket.MaximumBuysQuantity)
	}

	if time.Until(ticket.SalesEndDateTime) < time.Until(ticket.SalesStartDateTime) {
		return domain_errors.NewValidationError(
			"Date time to start ticket sales must be earlier then date time to end the sales",
			"SalesStartDateTime", ticket.SalesStartDateTime)
	}

	return nil

}
