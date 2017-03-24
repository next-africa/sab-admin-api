package countrystore

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"strings"
)

const COUNTRY_KIND = "Country"

type CountryNotFoundError struct {
	countryCode string
}

func (error CountryNotFoundError) Error() string {
	return fmt.Sprintf("Country with code %s not found", error.countryCode)
}

type Country struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func SaveCountry(country Country, gaeContext context.Context) error {
	key := datastore.NewKey(gaeContext, COUNTRY_KIND, strings.ToLower(country.Code), 0, nil)

	_, err := datastore.Put(gaeContext, key, &country)

	return err
}

func GetCountryByCode(code string, gaeContext context.Context) (Country, error) {
	key := datastore.NewKey(gaeContext, COUNTRY_KIND, strings.ToLower(code), 0, nil)

	var country Country

	err := datastore.Get(gaeContext, key, &country)

	if err == datastore.ErrNoSuchEntity {
		return country, &CountryNotFoundError{code}
	}

	return country, err
}
