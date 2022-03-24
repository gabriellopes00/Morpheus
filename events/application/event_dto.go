package application

// Params to create a event `Location` entity
type LocationParams struct {
	Street      string  `json:"street,omitempty"`
	District    string  `json:"district,omitempty"`
	State       string  `json:"state,omitempty"`
	City        string  `json:"city,omitempty"`
	Number      int     `json:"number,omitempty"`
	PostalCode  string  `json:"postal_code,omitempty"`
	Description string  `json:"description,omitempty"`
	Latitude    float64 `json:"latitude,omitempty"`
	Longitude   float64 `json:"longitude,omitempty"`
}

// Params to create a event `TicketOption` entity
type TicketParams struct {
	Title               string            `json:"title,omitempty"`
	Description         string            `json:"description,omitempty"`
	SalesStartDateTime  string            `json:"sales_start_datetime,omitempty"`
	SalesEndDateTime    string            `json:"sales_end_datetime,omitempty"`
	MinimumBuysQuantity int               `json:"maximum_buys_quantity,omitempty"`
	MaximumBuysQuantity int               `json:"minimum_buys_quantity,omitempty"`
	Lots                []TicketLotParams `json:"lots,omitempty"`
}

// Params to create a ticket option `Lot`
type TicketLotParams struct {
	Number   int     `json:"number,omitempty"`
	Price    float64 `json:"price,omitempty"`
	Quantity int     `json:"quantity,omitempty"`
}

// Params to create a `Event` entity
type CreateEventParams struct {
	Name               string         `json:"name,omitempty"`
	Description        string         `json:"description,omitempty"`
	CoverUrl           string         `json:"cover_url,omitempty"`
	OrganizerAccountId string         `json:"organizer_account_id,omitempty"`
	AgeGroup           int            `json:"age_group,omitempty"`
	Status             string         `json:"status,omitempty"`
	Location           LocationParams `json:"location" gorm:"foreignKey:EventId;references:Id"`
	Tickets            []TicketParams `json:"ticket_options,omitempty" gorm:"foreignKey:EventId;references:Id"`
	StartDateTime      string         `json:"start_datetime,omitempty"`
	EndDateTime        string         `json:"end_datetime,omitempty"`
	CategoryId         string         `json:"category,omitempty"`
	SubjectId          string         `json:"subject,omitempty"`
	Visibility         string         `json:"visibility,omitempty"`
}
