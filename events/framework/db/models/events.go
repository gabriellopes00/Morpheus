package models

import "time"

type EventModel struct {
	Id                 string `gorm:"primaryKey"`
	Name               string
	Description        string
	OrganizerAccountId string
	AgeGroup           int
	MaximumCapacity    int
	Status             string
	Location           EventLocationModel `gorm:"foreignKey:LocationId"`
	LocationId         string
	Duration           int
	TycketOptions      []TycketOptionModel
	Date               time.Time
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
