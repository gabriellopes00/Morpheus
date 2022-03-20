package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type TycketOption struct {
	Id        string      `json:"id,omitempty"`
	EventId   string      `json:"event_id,omitempty"`
	Title     string      `json:"title,omitempty"`
	Lots      []TycketLot `json:"lots,omitempty" gorm:"-"`
	CreatedAt time.Time   `json:"created_at,omitempty"`
}

func NewTycketOption(eventId, title string, lots []TycketLot) *TycketOption {
	return &TycketOption{
		Id:        uuid.NewV4().String(),
		EventId:   eventId,
		Title:     title,
		Lots:      lots,
		CreatedAt: time.Now(),
	}
}
