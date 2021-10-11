package entities

import (
	domain_errors "events/domain/errors"
	"fmt"
	"regexp"
	"strings"
)

type EventLocation struct {
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

func NewEventLocation(
	street, district, state, city, postalCode, description string,
	number int, latitude, longitude float64,
) (*EventLocation, error) {
	location := &EventLocation{
		Street:      strings.TrimSpace(street),
		District:    strings.TrimSpace(district),
		State:       strings.TrimSpace(state),
		City:        strings.TrimSpace(city),
		Number:      number,
		PostalCode:  strings.TrimSpace(postalCode),
		Description: strings.TrimSpace(description),
		Latitude:    latitude,
		Longitude:   longitude,
	}
	err := location.validate()
	if err != nil {
		return nil, err
	}

	return location, nil

}

func (l *EventLocation) validate() error {
	validLat, _ := regexp.MatchString(`^[-+]?([1-8]?\d(\.\d+)?|90(\.0+)?)$`, fmt.Sprintf("%f", l.Latitude))
	if !validLat {
		return domain_errors.NewValidationError("Invalid event location latitude format", "Latitude", l.Latitude)
	}

	validLng, _ := regexp.MatchString(`^[-+]?(180(\.0+)?|((1[0-7]\d)|([1-9]?\d))(\.\d+)?)$`, fmt.Sprintf("%f", l.Longitude))
	if !validLng {
		return domain_errors.NewValidationError("Invalid event location longitude format", "Longitude", l.Longitude)
	}

	validPostalCode, _ := regexp.MatchString(`^\d{5}[-]\d{3}$`, l.PostalCode)
	if !validPostalCode {
		return domain_errors.NewValidationError("Invalid event location postal code format", "PostalCode", l.PostalCode)
	}

	if l.Number <= 0 {
		return domain_errors.NewValidationError("Events' location number must be at least of 0", "Number", l.Number)
	}

	return nil
}
