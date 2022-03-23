package entities

// import (
// 	"time"
// )

// type EventVisibility string

// const (
// 	VisibilityPrivate     EventVisibility = "private"
// 	VisibilityPublic      EventVisibility = "private"
// 	VisibilityInvitedOnly EventVisibility = "invitedonly"
// )

// type A struct {
// 	Id                 string          `json:"id,omitempty" gorm:"primaryKey"`
// 	Name               string          `json:"name,omitempty"`
// 	Description        string          `json:"description,omitempty"` // rich text
// 	CoverUrl           string          `json:"cover_url,omitempty"`
// 	OrganizerAccountId string          `json:"organizer_account_id,omitempty"`
// 	AgeGroup           int             `json:"age_group,omitempty"`
// 	MaximumCapacity    int             `json:"maximum_capacity,omitempty"`
// 	Status             EventStatus     `json:"status,omitempty"`
// 	Location           EventLocation   `json:"location" gorm:"foreignKey:EventId;references:Id"`
// 	Tickets            []Ticket        `json:"ticket_options,omitempty" gorm:"foreignKey:EventId;references:Id"`
// 	StartDateTime      time.Time       `json:"start_datetime,omitempty"`
// 	EndDateTime        time.Time       `json:"end_datetime,omitempty"`
// 	CreatedAt          time.Time       `json:"created_at,omitempty"`
// 	UpdatedAt          time.Time       `json:"updated_at,omitempty"`
// 	CategoryId         string          `json:"category,omitempty"` // optional - category id
// 	SubjectId          string          `json:"subject,omitempty"`  // subject
// 	Visibility         EventVisibility `json:"visibility,omitempty"`
// }

// type Ticket_ struct {
// 	Title               string
// 	Description         string
// 	AllowSockTicket     bool
// 	SalesStartDateTime  time.Time `json:"sales_start_datetime,omitempty"`
// 	SalesEndDateTime    time.Time `json:"sales_end_datetime,omitempty"`
// 	MinimumBuysQuantity int
// 	MaximumBuysQuantity int
// 	Availability        string // either public or invited-only
// 	Lots                []TicketLot
// }

// type TicketLot_ struct {
// 	Number            int     `json:"number,omitempty"`
// 	Price             float64 `json:"price,omitempty"`
// 	AvailableQuantity int
// }
