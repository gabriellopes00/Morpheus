package domain_test

import (
	"events/domain/entities"
	"testing"
)

func TestNewEventLocation(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    entities.EventLocation
		expected error
	}{
		{entities.EventLocation{PostalCode: "00000-000", Number: 99}, nil},
		{entities.EventLocation{PostalCode: "00000-000", Number: 99}, nil},
		// {entities.EventLocation{PostalCode: "00000000"}, errors.New("")},
		// {entities.EventLocation{PostalCode: "00000-00"}, errors.New("")},
		// {entities.EventLocation{PostalCode: "0000000"}, errors.New("")},
		// {entities.EventLocation{PostalCode: "000000000"}, errors.New("")},
		// {entities.EventLocation{PostalCode: "000000-000"}, errors.New("")},
	}

	eventId := "07917a40-7d73-4b59-9c7d-7adf4ed18530"

	for _, test := range tests {
		_, actual := entities.NewEventLocation(
			test.param.Street, eventId, test.param.District, test.param.State, test.param.City, test.param.PostalCode, test.param.Description, test.param.Number, test.param.Latitude, test.param.Longitude)

		if actual != test.expected {
			t.Errorf("Expected NewEventLocation(%v) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}
