package graphql

import (
	"github.com/kr/pretty"
	"sab.com/domain/country"
	"sab.com/domain/university"
)

type InMemoryUniversityRepository struct {
	nextId       int64
	universities map[string]map[int64]university.University
}

func (r *InMemoryUniversityRepository) Save(univ *university.University, countryCode string) error {
	if !r.countryExists(countryCode) {
		r.universities[countryCode] = make(map[int64]university.University)
	}

	if univ.Id == 0 {
		r.nextId += 1
		univ.Id = r.nextId
	}
	r.universities[countryCode][univ.Id] = *univ
	return nil
}

func (r *InMemoryUniversityRepository) GetById(id int64, countryCode string) (university.University, error) {
	if !r.countryExists(countryCode) {
		return university.University{}, country.CountryNotFoundError
	}
	if univ, ok := r.universities[countryCode][id]; ok {
		return univ, nil
	} else {
		return university.University{}, university.UniversityNotFoundError
	}

}

func (r *InMemoryUniversityRepository) GetAll(countryCode string) ([]university.University, error) {
	var universities []university.University
	if !r.countryExists(countryCode) {
		return universities, nil
	}
	universityMap, _ := r.universities[countryCode]
	universities = make([]university.University, len(universityMap))
	for _, univ := range universityMap {
		universities = append(universities, univ)
	}
	return universities, nil
}

func (r *InMemoryUniversityRepository) HasUniversity(id int64, countryCode string) (bool, error) {
	if !r.countryExists(countryCode) {
		return false, nil
	}
	ctr := r.universities[countryCode]
	_, universityExists := ctr[id]

	return universityExists, nil
}

func (r *InMemoryUniversityRepository) countryExists(countryCode string) (result bool) {
	_, result = r.universities[countryCode]
	return
}

type InMemoryCountryRepository struct {
	countries map[string]country.Country
}

func (r *InMemoryCountryRepository) Save(ctr *country.Country) error {
	if r.countries == nil {
		r.countries = make(map[string]country.Country, 1)
	}
	r.countries[ctr.Code] = *ctr
	return nil
}

func (r *InMemoryCountryRepository) GetByCode(code string) (country.Country, error) {
	if ctr, ok := r.countries[code]; ok {
		return ctr, nil
	} else {
		return country.Country{}, country.CountryNotFoundError
	}
}

func (r *InMemoryCountryRepository) GetAll() ([]country.Country, error) {
	countries := make([]country.Country, len(r.countries))
	for _, aCountry := range r.countries {
		countries = append(countries, aCountry)
	}
	return countries, nil
}

func (r *InMemoryCountryRepository) HasCountryWithCode(code string) (result bool, err error) {
	_, result = r.countries[code]
	return
}

func Diff(a, b interface{}) []string {
	return pretty.Diff(a, b)
}
