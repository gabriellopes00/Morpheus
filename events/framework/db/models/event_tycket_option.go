package models

import "time"

type TycketOptionModel struct {
	Id        string `gorm:"primaryKey"`
	EventId   string
	Title     string
	Lots      []TycketLotModel
	CreatedAt time.Time
}
