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

// Params to create a event `TycketOption` entity
type TycketOptionsParams struct {
	Title string             `json:"title,omitempty"`
	Lots  []TycketLotsParams `json:"lots,omitempty"`
}

// Params to create a tycket option `Lot`
type TycketLotsParams struct {
	Number       int     `json:"number,omitempty"`
	TycketPrice  float64 `json:"tycket_price,omitempty"`
	TycketAmount int     `json:"tycket_amount,omitempty"`
}

// Params to create a `Event` entity
type CreateEventParams struct {
	Name               string                `json:"name,omitempty"`
	Description        string                `json:"description,omitempty"`
	OrganizerAccountId string                `json:"organizer_account_id,omitempty"`
	AgeGroup           int                   `json:"age_group,omitempty"`
	MaximumCapacity    int                   `json:"maximum_capacity,omitempty"`
	Location           LocationParams        `json:"location,omitempty"`
	TycketOptions      []TycketOptionsParams `json:"tycket_options,omitempty"`
	Date               string                `json:"date,omitempty"`
	Duration           int                   `json:"duration,omitempty"`
}
