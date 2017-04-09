package country

type CountryService struct {
	countryRepository CountryRepository
}

func NewCountryService(countryRepository CountryRepository) CountryService {
	return CountryService{countryRepository}
}

func (service *CountryService) SaveCountry(country *Country) error {
	return service.countryRepository.Save(country)
}

func (service *CountryService) GetCountryByCode(countryCode string) (Country, error) {
	return service.countryRepository.GetByCode(countryCode)
}

func (service *CountryService) GetAllCountries() ([]Country, error) {
	return service.countryRepository.GetAll()
}
