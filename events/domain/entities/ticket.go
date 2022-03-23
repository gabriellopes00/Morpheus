package entities

import (
	"events/domain"
	"time"
)

type TicketAvailability string

const (
	AvailabilityPrivate     TicketAvailability = "private"
	AvailabilityPublic      TicketAvailability = "public"
	AvailabilityInvitedOnly TicketAvailability = "invited_only"
)

type Ticket struct {
	domain.Entity
	Title               string             `json:"title,omitempty"`
	Description         string             `json:"description,omitempty"`
	SalesStartDateTime  time.Time          `json:"sales_start_datetime,omitempty"`
	SalesEndDateTime    time.Time          `json:"sales_end_datetime,omitempty"`
	MinimumBuysQuantity int                `json:"maximum_buys_quantity,omitempty"`
	MaximumBuysQuantity int                `json:"minimum_buys_quantity,omitempty"`
	Availability        TicketAvailability `json:"availability,omitempty"`
	Lots                []TicketLot        `json:"lots,omitempty"`
}
