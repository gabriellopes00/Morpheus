package models

import "time"

type TycketLotModel struct {
	Id             string `gorm:"primaryKey"`
	Number         int
	TycketOptionId string
	TycketPrice    float64
	TycketAmount   int
	CreatedAt      time.Time
}
