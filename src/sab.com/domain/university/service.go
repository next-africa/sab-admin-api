package university

import "sab.com/domain/country"

type UniversityService struct {
	universityRepository UniversityRepository
	countryRepository    country.CountryRepository
}

func NewUniversityService(universityRepository UniversityRepository, countryRepository country.CountryRepository) UniversityService {
	return UniversityService{universityRepository, countryRepository}
}

func (service *UniversityService) CreateUniversity(theUniversity *University, countryCode string) error {
	if countryWithCodeExists, err := service.countryRepository.HasCountryWithCode(countryCode); err != nil {
		return err
	} else {
		if countryWithCodeExists {
			return service.universityRepository.Save(theUniversity, countryCode)
		} else {
			return country.CountryNotFoundError
		}
	}
}

func (service *UniversityService) UpdateUniversity(theUniversity *University, countryCode string) error {
	if universityExists, err := service.universityRepository.HasUniversity(theUniversity.Id, countryCode); err != nil {
		return err
	} else {
		if universityExists {
			return service.universityRepository.Save(theUniversity, countryCode)
		} else {
			return UniversityNotFoundError
		}
	}
}

func (service *UniversityService) GetUniversityByIdAndCountryCode(id int64, countryCode string) (University, error) {
	return service.universityRepository.GetById(id, countryCode)
}

func (service *UniversityService) GetAllUniversitiesForCountryCode(countryCode string) ([]University, error) {
	return service.universityRepository.GetAll(countryCode)
}
