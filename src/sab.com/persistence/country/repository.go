package country

import (
	"google.golang.org/appengine/datastore"
	"sab.com/domain/country"
	"sab.com/persistence"
	"strings"
)

const COUNTRY_KIND = "Country"

type DatastoreCountryRepository struct {
	contextStore *persistence.ContextStore
}

func NewCountryRepository(contextStore *persistence.ContextStore) DatastoreCountryRepository {
	return DatastoreCountryRepository{contextStore}
}

func (repository DatastoreCountryRepository) Save(countryToSave *country.Country) error {
	gaeContext := repository.contextStore.GetContext()
	key := datastore.NewKey(gaeContext, COUNTRY_KIND, strings.ToLower(countryToSave.Code), 0, nil)

	key, err := datastore.Put(gaeContext, key, countryToSave)

	return err
}

func (repository DatastoreCountryRepository) GetByCode(code string) (country.Country, error) {
	gaeContext := repository.contextStore.GetContext()

	key := datastore.NewKey(gaeContext, COUNTRY_KIND, strings.ToLower(code), 0, nil)

	var countryToReturn country.Country

	err := datastore.Get(gaeContext, key, &countryToReturn)

	if err == datastore.ErrNoSuchEntity {
		return countryToReturn, country.CountryNotFoundError
	}

	return countryToReturn, err
}

func (repository DatastoreCountryRepository) GetAll() ([]country.Country, error) {
	countries := make([]country.Country, 0)

	_, err := datastore.NewQuery(COUNTRY_KIND).GetAll(repository.contextStore.GetContext(), &countries)

	if err != nil {
		return countries, err
	}

	return countries, nil
}

func (repository DatastoreCountryRepository) HasCountryWithCode(code string) (bool, error) {
	gaeContext := repository.contextStore.GetContext()

	countryKey := datastore.NewKey(gaeContext, COUNTRY_KIND, strings.ToLower(code), 0, nil)

	var dst []country.Country

	q, err := datastore.NewQuery(COUNTRY_KIND).Filter("__key__ =", countryKey).KeysOnly().GetAll(gaeContext, dst)

	return q != nil, err
}
