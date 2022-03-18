package application

import (
	"events/domain/entities"
	"events/framework/db/repositories"
)

type createEventUsecase struct {
	repository repositories.EventsRepository
}

func NewCreateEventUsecase(repo repositories.EventsRepository) *createEventUsecase {
	return &createEventUsecase{
		repository: repo,
	}
}

type CreateEventParams struct {
	Name               string `json:"name,omitempty"`
	Description        string `json:"description,omitempty"`
	OrganizerAccountId string `json:"organizer_account_id,omitempty"`
	AgeGroup           int    `json:"age_group,omitempty"`
	MaximumCapacity    int    `json:"maximum_capacity,omitempty"`
	Location           struct {
		Street      string  `json:"street,omitempty"`
		District    string  `json:"district,omitempty"`
		State       string  `json:"state,omitempty"`
		City        string  `json:"city,omitempty"`
		Number      int     `json:"number,omitempty"`
		PostalCode  string  `json:"postal_code,omitempty"`
		Description string  `json:"description,omitempty"`
		Latitude    float64 `json:"latitude,omitempty"`
		Longitude   float64 `json:"longitude,omitempty"`
	} `json:"location,omitempty"`
	TycketOptions []struct {
		Title string `json:"title,omitempty"`
		Lots  []struct {
			Number       int     `json:"number,omitempty"`
			TycketPrice  float64 `json:"tycket_price,omitempty"`
			TycketAmount int     `json:"tycket_amount,omitempty"`
		} `json:"lots,omitempty"`
	} `json:"tycket_options,omitempty"`
	Date     string `json:"date,omitempty"`
	Duration int    `json:"duration,omitempty"`
}

func (c *createEventUsecase) Create(params *CreateEventParams) (*entities.Event, error) {

	location := params.Location
	eventLocation, err := entities.NewEventLocation(
		location.Street, location.District, location.State, location.City,
		location.PostalCode, location.Description, location.Number,
		location.Latitude, params.Location.Longitude)
	if err != nil {
		return nil, err
	}

	event, err := entities.NewEvent(
		params.Name, params.Description, params.OrganizerAccountId, params.AgeGroup,
		params.MaximumCapacity, eventLocation, params.Duration, nil, params.Date)
	if err != nil {
		return nil, err
	}

	options := params.TycketOptions
	var tycketOptions []entities.TycketOption

	for _, option := range options {

		tycketOption := entities.NewTycketOption(event.Id, option.Title, nil)

		lots := option.Lots
		var tycketLots []entities.TycketLot

		for _, lot := range lots {
			tycketLot := entities.NewTycketLot(lot.Number, tycketOption.Id, lot.TycketPrice, lot.TycketAmount)
			tycketLots = append(tycketLots, *tycketLot)
		}

		tycketOption.Lots = tycketLots
		tycketOptions = append(tycketOptions, *tycketOption)
	}

	event.TycketOptions = tycketOptions

	// if err = c.repository.Create(event); err != nil {
	// 	return nil, err
	// }

	return event, nil
}
