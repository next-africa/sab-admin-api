package country

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

const (
	COUNTRY_CODE  = "CA"
	COUNTRY_NAME  = "Canada"
	COUNTRY_CODE1 = "US"
	COUNTRY_NAME1 = "United States"
)

type CountryServiceTestSuite struct {
	suite.Suite
	countryRepository *CountryRepositoryMock
	service           CountryService
}

func (suite *CountryServiceTestSuite) SetupTest() {
	suite.countryRepository = new(CountryRepositoryMock)
	suite.service = NewCountryService(suite.countryRepository)
}

func (suite *CountryServiceTestSuite) TestSaveCountry() {
	aCountry := Country{Code: COUNTRY_CODE, Name: COUNTRY_NAME}
	suite.countryRepository.On("Save", &aCountry).Return(nil)

	suite.service.SaveCountry(&aCountry)

	suite.countryRepository.AssertExpectations(suite.T())
}

func (suite *CountryServiceTestSuite) TestGetByCoundryCode() {
	expectedCountry := Country{Code: COUNTRY_CODE, Name: COUNTRY_NAME}
	suite.countryRepository.On("GetByCode", COUNTRY_CODE).Return(expectedCountry, nil)

	retrievedCountry, err := suite.service.GetCountryByCode(COUNTRY_CODE)

	suite.NoError(err)
	suite.Equal(expectedCountry, retrievedCountry)
}

func (suite *CountryServiceTestSuite) TestGetAllCountries() {
	country1 := Country{Code: COUNTRY_CODE, Name: COUNTRY_NAME}
	country2 := Country{Code: COUNTRY_CODE1, Name: COUNTRY_NAME1}
	suite.countryRepository.On("GetAll").Return([]Country{country1, country2}, nil)

	countries, err := suite.service.GetAllCountries()

	suite.NoError(err)
	suite.Contains(countries, country1, country2)
}

func TestUniversityService(t *testing.T) {
	suite.Run(t, new(CountryServiceTestSuite))
}
