package application

import (
	"encoding/json"
	"errors"
	"events/domain/entities"
	"events/framework/db/repositories"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type FindEvents struct {
	repository repositories.EventsRepository
}

func NewFindEvents(repo repositories.EventsRepository) *FindEvents {
	return &FindEvents{
		repository: repo,
	}
}

func (u *FindEvents) FindAccountEvents(accountId string) ([]entities.Event, error) {
	return u.repository.FindAccountEvents(accountId)
}

func (u *FindEvents) FindEventById(eventId string) (*entities.Event, error) {
	return u.repository.FindById(eventId)
}

func (u *FindEvents) FindNearEvents(latitude, longitude float64) ([]entities.Event, error) {
	// request https://api.bigdatacloud.net/data/reverse-geocode-client?latitude=37.42159&longitude=-122.0837&localityLanguage=en
	// get city name
	// fetch events by this city name ==> cache it

	url := fmt.Sprintf("https://api.bigdatacloud.net/data/reverse-geocode-client?latitude=%flongitude=%f&localityLanguage=en", latitude, longitude)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	type Response struct {
		Latitude                 string `json:"latitude"`
		Longitude                string `json:"longitude"`
		Continent                string `json:"continent"`
		LookupSource             string `json:"lookupSource"`
		ContinentCode            string `json:"continentCode"`
		City                     string `json:"city"`
		CountryName              string `json:"countryName"`
		PostCode                 string `json:"postcode"`
		CountryCode              string `json:"countryCode"`
		PrincipalSubdivision     string `json:"principalSubdivision"`
		PrincipalSubdivisionCode string `json:"principalSubdivisionCode"`
	}

	var data Response
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	fmt.Println(data)

	return nil, nil
}

func (u *FindEvents) FindAll(state string, month, ageGroup, limit, offset int) ([]entities.Event, error) {
	if len(state) != 2 {
		return nil, errors.New("invalid state abbreviation")
	}

	if month < 0 || month > 12 {
		return nil, errors.New("invalid month")
	}

	if limit < 1 && limit > 30 {
		return nil, errors.New("invalid results limit")
	}

	switch ageGroup {
	case 0, 10, 12, 14, 16, 18:
		break
	default:
		return nil, errors.New("age group must be: 0, 10, 12, 14, 16 or 18")

	}

	return u.repository.FindAll(state, month, ageGroup, limit, offset)
}
