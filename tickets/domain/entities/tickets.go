package entities

import (
	"time"

	gouuid "github.com/satori/go.uuid"
)

type Ticket struct {
	Id             string    `json:"id,omitempty"`
	EventId        string    `json:"event_id,omitempty"`
	OwnerAccountId string    `json:"owner_account_id,omitempty"`
	ExpirationDate time.Time `json:"expiration_date,omitempty"`
	UsedAt         time.Time `json:"used_at,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}

func NewTicket(eventId, ownerAccountId string, expirationDate time.Time) *Ticket {
	return &Ticket{
		Id:             gouuid.NewV4().String(),
		EventId:        eventId,
		OwnerAccountId: ownerAccountId,
		ExpirationDate: expirationDate,
		CreatedAt:      time.Now().Local(),
	}
}
