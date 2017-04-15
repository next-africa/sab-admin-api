package university

import (
	"errors"
)

var (
	UniversityNotFoundError = errors.New("Country with the given ID was not found")
)

type UniversityRepository interface {
	Save(univ *University, countryCode string) error
	GetById(id int64, countryCode string) (University, error)
	GetAll(countryCode string) ([]University, error)
}
