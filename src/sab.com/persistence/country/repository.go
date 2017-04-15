package country

import (
	"google.golang.org/appengine/datastore"
	"sab.com/domain/country"
	"sab.com/persistence"
	"strings"
)

const COUNTRY_KIND = "Country"

type countryRepository struct {
	contextStore *persistence.ContextStore
}

func NewCountryRepository(contextStore *persistence.ContextStore) countryRepository {
	return countryRepository{contextStore}
}

func (repository countryRepository) Save(countryToSave *country.Country) error {
	gaeContext := repository.contextStore.GetContext()
	key := datastore.NewKey(gaeContext, COUNTRY_KIND, strings.ToLower(countryToSave.Code), 0, nil)

	key, err := datastore.Put(gaeContext, key, countryToSave)

	return err
}

func (repository countryRepository) GetByCode(code string) (country.Country, error) {
	gaeContext := repository.contextStore.GetContext()

	key := datastore.NewKey(gaeContext, COUNTRY_KIND, strings.ToLower(code), 0, nil)

	var countryToReturn country.Country

	err := datastore.Get(gaeContext, key, &countryToReturn)

	if err == datastore.ErrNoSuchEntity {
		return countryToReturn, country.CountryNotFoundError
	}

	return countryToReturn, err
}

func (repository countryRepository) GetAll() ([]country.Country, error) {
	countries := make([]country.Country, 0)

	_, err := datastore.NewQuery(COUNTRY_KIND).GetAll(repository.contextStore.GetContext(), &countries)

	if err != nil {
		return countries, err
	}

	return countries, nil
}

func (repository countryRepository) HasCountryWithCode(code string) (bool, error) {
	if _, err := repository.GetByCode(code); err != nil {
		if err == country.CountryNotFoundError {
			return false, nil
		} else {
			return false, err
		}
	}

	return true, nil

}
