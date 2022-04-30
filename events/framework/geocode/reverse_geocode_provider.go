package geocode

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type LocationData struct {
	Latitude                 float64 `json:"latitude"`
	Longitude                float64 `json:"longitude"`
	Continent                string  `json:"continent"`
	LookupSource             string  `json:"lookupSource"`
	ContinentCode            string  `json:"continentCode"`
	City                     string  `json:"city"`
	CountryName              string  `json:"countryName"`
	PostCode                 string  `json:"postcode"`
	CountryCode              string  `json:"countryCode"`
	PrincipalSubdivision     string  `json:"principalSubdivision"`
	PrincipalSubdivisionCode string  `json:"principalSubdivisionCode"`
}

type GeocodeProvider struct{}

func NewGeocodeProvider() *GeocodeProvider {
	return &GeocodeProvider{}
}

func (*GeocodeProvider) Reverse(latitude, longitude string) (*LocationData, error) {
	url := fmt.Sprintf("https://api.bigdatacloud.net/data/reverse-geocode-client?latitude=%s&longitude=%s&localityLanguage=en", latitude, longitude)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data LocationData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
