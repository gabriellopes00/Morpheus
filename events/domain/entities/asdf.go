package entities

import (
	"time"
)

type A struct {
	Id                 string        `json:"id,omitempty" gorm:"primaryKey"`
	Name               string        `json:"name,omitempty"`
	Description        string        `json:"description,omitempty"` // html
	OrganizerAccountId string        `json:"organizer_account_id,omitempty"`
	AgeGroup           int           `json:"age_group,omitempty"`
	MaximumCapacity    int           `json:"maximum_capacity,omitempty"`
	Status             EventStatus   `json:"status,omitempty"`
	Location           EventLocation `json:"location" gorm:"foreignKey:EventId;references:Id"`
	Duration           int           `json:"duration,omitempty"`
	Tyckets            []Tycket      `json:"tycket_options,omitempty" gorm:"foreignKey:EventId;references:Id"`
	StartDateTime      time.Time     `json:"start_datetime,omitempty"`
	EndDateTime        time.Time     `json:"end_datetime,omitempty"`
	CreatedAt          time.Time     `json:"created_at,omitempty"`
	UpdatedAt          time.Time     `json:"updated_at,omitempty"`
	Category           string        `` // optional
	Subject            string        // subject id
	Visibility         string        // either private, public or invited-only
}

/*

free ingress

*/

type Tycket struct {
	Title               string
	Description         string
	AllowSockTycket     bool
	SalesStartDateTime  time.Time `json:"sales_start_datetime,omitempty"`
	SalesEndDateTime    time.Time `json:"sales_end_datetime,omitempty"`
	MinimumBuysQuantity int
	MaximumBuysQuantity int
	Availability        string // either public or invited-only
	Lots                []struct {
		Number            int     `json:"number,omitempty"`
		Price             float64 `json:"price,omitempty"`
		AvailableQuantity int
	}
}
