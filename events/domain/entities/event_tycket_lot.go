package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type TycketLot struct {
	Id             string    `json:"id,omitempty"`
	Number         int       `json:"number,omitempty"`
	TycketOptionId string    `json:"tycket_option_id,omitempty" gorm:"column:event_tycket_option_id"`
	TycketPrice    float64   `json:"tycket_price,omitempty"`
	TycketAmount   int       `json:"tycket_amount,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}

func NewTycketLot(number int, tycketOptionId string, tycketPrice float64, tycketAmount int) *TycketLot {
	return &TycketLot{
		Id:             uuid.NewV4().String(),
		Number:         number,
		TycketOptionId: tycketOptionId,
		TycketPrice:    tycketPrice, // TODO: validate positive numbers
		TycketAmount:   tycketAmount,
		CreatedAt:      time.Now(),
	}
}
