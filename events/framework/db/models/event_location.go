package models

import "time"

type EventLocationModel struct {
	Id          string `gorm:"primaryKey"`
	EventId     string
	Street      string
	District    string
	State       string
	City        string
	Number      int
	PostalCode  string
	Description string
	Latitude    float64
	Longitude   float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
