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

type CountryGraphService struct {
	countryService *country.CountryService
}

func NewCountryGraphqlService(countryService *country.CountryService) CountryGraphService {
	return CountryGraphService{countryService: countryService}
}

func (graphqlService *CountryGraphService) mapCountriesToCountryNodes(countries []country.Country) []CountryNode {
	countriesMap := make([]CountryNode, len(countries))
	for i, v := range countries {
		countriesMap[i] = graphqlService.NewCountryNodeFromCountry(&v)
	}
	return countriesMap
}

func (graphqlService *CountryGraphService) NewCountryNodeFromCountry(aCountry *country.Country) CountryNode {
	id := fmt.Sprintf("%s:%s", "Country", aCountry.Code)
	id = base64.StdEncoding.EncodeToString([]byte(id))

	return CountryNode{id, aCountry}
}

func (graphqlService *CountryGraphService) GetCountryByGlobalId(encodedGlobalId string) (CountryNode, error) {
	if decoded, err := base64.StdEncoding.DecodeString(encodedGlobalId); err != nil {
		return CountryNode{}, err
	} else {
		idParts := strings.Split(string(decoded), ":")

		if len(idParts) != 2 {
			return CountryNode{}, errors.New("Invalid global country Id, the country relay Id should be of the form Country:{countryCode}")
		}

		code := idParts[1]

		return graphqlService.GetCountryNodeByCode(code)
	}
}

func (graphqlService *CountryGraphService) GetCountryNodeByCode(code string) (CountryNode, error) {
	theCountry, err := graphqlService.countryService.GetCountryByCode(code)

	if err != nil {
		return CountryNode{}, err
	}
	return graphqlService.NewCountryNodeFromCountry(&theCountry), nil
}

func (graphqlService *CountryGraphService) GetAllCountries() ([]CountryNode, error) {
	if countries, err := graphqlService.countryService.GetAllCountries(); err != nil {
		return []CountryNode{}, nil
	} else {
		return graphqlService.mapCountriesToCountryNodes(countries), nil
	}
}

func (graphqlService *CountryGraphService) SaveCountry(country *country.Country) error {
	return graphqlService.countryService.SaveCountry(country)
}
