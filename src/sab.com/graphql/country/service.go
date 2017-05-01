package graphql

import (
	"sab.com/domain/country"
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
	for i := range countries {
		countriesMap[i] = graphqlService.NewCountryNodeFromCountry(&countries[i])
	}
	return countriesMap
}

func (graphqlService *CountryGraphService) NewCountryNodeFromCountry(aCountry *country.Country) CountryNode {
	return CountryNode{aCountry.Code, aCountry}
}

func (graphqlService *CountryGraphService) GetCountryByGlobalId(code string) (CountryNode, error) {
	return graphqlService.GetCountryNodeByCode(code)
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
