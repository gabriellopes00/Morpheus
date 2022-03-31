package entities

import (
	"events/domain"
	domain_errors "events/domain/errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type TicketOption struct {
	domain.Entity
	Title               string            `json:"title,omitempty"`
	Description         string            `json:"description,omitempty"`
	SalesStartDateTime  time.Time         `json:"sales_start_datetime,omitempty"`
	SalesEndDateTime    time.Time         `json:"sales_end_datetime,omitempty"`
	MinimumBuysQuantity int               `json:"maximum_buys_quantity,omitempty"`
	MaximumBuysQuantity int               `json:"minimum_buys_quantity,omitempty"`
	Lots                []TicketOptionLot `json:"lots,omitempty"`
}

func NewTicketOption(
	title, description, salesStartDateTime, salesEndDateTime string,
	maximumBuysQuantity, minimumBuysQuantity int, lots []TicketOptionLot) (*TicketOption, error) {

	var err error

	ticketOption := new(TicketOption)

	ticketOption.Id = uuid.NewV4().String()
	ticketOption.CreatedAt = time.Now()
	ticketOption.UpdatedAt = time.Now()

	ticketOption.Title = title
	ticketOption.Description = description
	ticketOption.MaximumBuysQuantity = maximumBuysQuantity
	ticketOption.MinimumBuysQuantity = minimumBuysQuantity
	ticketOption.Lots = lots

	ticketOption.SalesStartDateTime, err = time.Parse(time.RFC3339, salesStartDateTime)
	if err != nil {
		return nil, domain_errors.NewValidationError(
			`TicketOption's dates must be in "RFC3339" format`,
			"SalesStartDateTime", salesStartDateTime)
	}

	ticketOption.SalesEndDateTime, err = time.Parse(time.RFC3339, salesEndDateTime)
	if err != nil {
		return nil, domain_errors.NewValidationError(
			`TicketOption's dates must be in "RFC3339" format`,
			"SalesEndDateTime", salesEndDateTime)
	}

	err = ticketOption.validate()
	if err != nil {
		return nil, err
	}

	return ticketOption, nil
}

func (ticketOption *TicketOption) validate() error {
	if len(ticketOption.Title) < 1 || len(ticketOption.Title) > 255 {
		return domain_errors.NewValidationError(
			"TicketOption's title must have at least of 4 characters and at most of 255",
			"title", ticketOption.Title)
	}

	if ticketOption.MinimumBuysQuantity < 1 {
		return domain_errors.NewValidationError(
			"Minimum buys ticketOptions must be greather than 1",
			"MinimumBuysQuantity", ticketOption.MinimumBuysQuantity)
	}

	if ticketOption.MaximumBuysQuantity < ticketOption.MinimumBuysQuantity {
		return domain_errors.NewValidationError(
			"Maximum buys ticketOptions must be greather than Minimum buys ticketOptions",
			"MaximumBuysQuantity", ticketOption.MaximumBuysQuantity)
	}

	if time.Until(ticketOption.SalesEndDateTime) < time.Until(ticketOption.SalesStartDateTime) {
		return domain_errors.NewValidationError(
			"Date time to start ticketOption sales must be earlier then date time to end the sales",
			"SalesStartDateTime", ticketOption.SalesStartDateTime)
	}

	return nil

}
