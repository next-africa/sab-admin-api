package country

import (
	"errors"
)

var (
	CountryNotFoundError = errors.New("The country with the given code was not found")
)

type CountryRepository interface {
	Save(ctr *Country) error
	GetByCode(code string) (Country, error)
	GetAll() ([]Country, error)
	HasCountryWithCode(code string) (bool, error)
}
