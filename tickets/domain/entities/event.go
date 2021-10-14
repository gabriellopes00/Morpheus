package entities

import "time"

type Event struct {
	Id   string    `json:"id,omitempty"`
	Date time.Time `json:"date,omitempty"`
}

func NewEvent(id string, date time.Time) *Event {
	return &Event{
		Id:   id,
		Date: date,
	}
}
