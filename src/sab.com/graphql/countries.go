package graphql

import (
	"encoding/base64"
	"errors"
	"fmt"
	"sab.com/domain/country"
	"strings"
)

type CountryNode struct {
	Id         string           `json:"id"`
	Properties *country.Country `json:"properties"`
}

func newCountryNodeFromCountry(aCountry *country.Country) CountryNode {
	id := fmt.Sprintf("%s:%s", "Country", aCountry.Code)
	id = base64.StdEncoding.EncodeToString([]byte(id))

	return CountryNode{id, aCountry}
}

func mapCountriesToCountryNodes(countries []country.Country) []CountryNode {
	countriesMap := make([]CountryNode, len(countries))
	for i, v := range countries {
		countriesMap[i] = newCountryNodeFromCountry(&v)
	}
	return countriesMap
}

func getCountryByGlobalId(encodedGlobalId string, countryService *country.CountryService) (CountryNode, error) {
	if decoded, err := base64.StdEncoding.DecodeString(encodedGlobalId); err != nil {
		return CountryNode{}, err
	} else {
		idParts := strings.Split(string(decoded), ":")

		if len(idParts) != 2 {
			return CountryNode{}, errors.New("Invalid global country Id, the country relay Id should be of the form Country:{countryCode}")
		}

		code := idParts[1]

		return getCountryNodeByCode(code, countryService)
	}
}

func getCountryNodeByCode(code string, countryService *country.CountryService) (CountryNode, error) {
	theCountry, err := countryService.GetCountryByCode(code)

	if err != nil {
		return CountryNode{}, err
	}

	return newCountryNodeFromCountry(&theCountry), nil
}

func getAllCountries(countryService *country.CountryService) ([]CountryNode, error) {
	if countries, err := countryService.GetAllCountries(); err != nil {
		return []CountryNode{}, nil
	} else {
		return mapCountriesToCountryNodes(countries), nil
	}
}
