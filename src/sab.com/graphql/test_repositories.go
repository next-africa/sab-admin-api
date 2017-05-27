package graphql

import (
	"github.com/kr/pretty"
	"sab.com/domain/country"
	"sab.com/domain/university"
)

type InMemoryCountryRepository struct {
	countries []country.Country
}

func (r *InMemoryCountryRepository) Save(ctr *country.Country) error {
	r.countries = append(r.countries, *ctr)
	return nil
}

func (r *InMemoryCountryRepository) GetByCode(code string) (country.Country, error) {
	for _, ctr := range r.countries {
		if ctr.Code == code {
			return ctr, nil
		}
	}
	return country.Country{}, country.CountryNotFoundError
}

func (r *InMemoryCountryRepository) GetAll() ([]country.Country, error) {
	return r.countries, nil
}

func (r *InMemoryCountryRepository) HasCountryWithCode(code string) (result bool, err error) {
	_, err = r.GetByCode(code)
	result = err == nil
	err = nil
	return
}

type InMemoryUniversityRepository struct {
	nextId       int64
	universities map[string][]university.University
}

func (r *InMemoryUniversityRepository) Save(univ *university.University, countryCode string) error {
	if r.universities == nil {
		r.universities = make(map[string][]university.University)
	}
	if !r.countryExists(countryCode) {
		r.universities[countryCode] = make([]university.University, 0)
	}

	if univ.Id == 0 {
		r.nextId += 1
		univ.Id = r.nextId
	}
	r.universities[countryCode] = append(r.universities[countryCode], *univ)
	return nil
}

func (r *InMemoryUniversityRepository) GetById(id int64, countryCode string) (university.University, error) {
	if !r.countryExists(countryCode) {
		return university.University{}, university.UniversityNotFoundError
	}
	for _, univ := range r.universities[countryCode] {
		if univ.Id == id {
			return univ, nil
		}
	}

	return university.University{}, university.UniversityNotFoundError
}

func (r *InMemoryUniversityRepository) GetAll(countryCode string) ([]university.University, error) {
	if !r.countryExists(countryCode) {
		return make([]university.University, 0), nil
	}
	return r.universities[countryCode], nil
}

func (r *InMemoryUniversityRepository) HasUniversity(id int64, countryCode string) (result bool, err error) {
	_, err = r.GetById(id, countryCode)
	result = err == nil
	err = nil
	return
}

func (r *InMemoryUniversityRepository) countryExists(countryCode string) (result bool) {
	_, result = r.universities[countryCode]
	return
}

func Diff(a, b interface{}) []string {
	return pretty.Diff(a, b)
}
