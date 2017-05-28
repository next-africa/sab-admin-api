package admin_module

import (
	"sab.com/domain/country"
	"sab.com/domain/university"
	"sab.com/persistence"
	countryPersistence "sab.com/persistence/country"
	universityPersistence "sab.com/persistence/university"
)

var (
	contextStore persistence.ContextStore

	countryRepository country.CountryRepository = nil
	countryService    country.CountryService

	universityRepository university.UniversityRepository = nil
	universityService    university.UniversityService
)

func GetContextStore() *persistence.ContextStore {
	return &contextStore
}

func GetCountryService() *country.CountryService {
	if (country.CountryService{}) == countryService {
		countryService = country.NewCountryService(getCountryRepository())
	}
	return &countryService
}

func GetUniversityService() *university.UniversityService {
	if (university.UniversityService{}) == universityService {
		universityService = university.NewUniversityService(getUniversityRepository(), getCountryRepository())
	}
	return &universityService
}

func getCountryRepository() country.CountryRepository {
	if countryRepository == nil {
		countryRepository = countryPersistence.NewCountryRepository(GetContextStore())
	}
	return countryRepository
}

func getUniversityRepository() university.UniversityRepository {
	if universityRepository == nil {
		universityRepository = universityPersistence.NewUniversityRepository(GetContextStore())
	}
	return universityRepository
}
