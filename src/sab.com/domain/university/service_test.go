package university

import (
	"github.com/stretchr/testify/suite"
	"sab.com/domain/country"
	"testing"
)

const (
	VALID_COUNTRY_CODE         = "CA"
	INVALID_COUNTRY_CODE       = "US"
	UNIVERSITY_ID        int64 = 123456789.
)

type UniversityServiceTestSuite struct {
	suite.Suite
	universityRepository *UniversityRepositoryMock
	countryRepository    *country.CountryRepositoryMock
	service              UniversityService
}

func (suite *UniversityServiceTestSuite) SetupTest() {
	suite.universityRepository = new(UniversityRepositoryMock)
	suite.countryRepository = new(country.CountryRepositoryMock)
	suite.service = NewUniversityService(suite.universityRepository, suite.countryRepository)
}

func (suite *UniversityServiceTestSuite) TestSaveUniversityGivenUniversityAndValidCountryCode() {
	suite.countryRepository.On("HasCountryWithCode", VALID_COUNTRY_CODE).Return(true, nil)
	university := new(University)
	suite.universityRepository.On("Save", university, VALID_COUNTRY_CODE).Return(nil)

	if err := suite.service.SaveUniversity(university, VALID_COUNTRY_CODE); err != nil {
		suite.Fail("Save University Failed", err)
	}

	suite.universityRepository.AssertExpectations(suite.T())
}

func (suite *UniversityServiceTestSuite) TestSaveUniversityGivenUniversityAndInvalidCountryCode() {
	suite.countryRepository.On("HasCountryWithCode", INVALID_COUNTRY_CODE).Return(false, nil)
	university := new(University)

	err := suite.service.SaveUniversity(university, INVALID_COUNTRY_CODE)

	if suite.Error(err) {
		suite.Equal(country.CountryNotFoundError, err)
	}
}

func (suite *UniversityServiceTestSuite) TestGetUniversityByIdAndCountryCodeGivenUniversityDoesNotExists() {
	suite.universityRepository.On("GetById", UNIVERSITY_ID, VALID_COUNTRY_CODE).Return(*new(University), UniversityNotFoundError)

	_, err := suite.service.getUniversityByIdAndCountryCode(UNIVERSITY_ID, VALID_COUNTRY_CODE)

	if suite.Error(err) {
		suite.Equal(UniversityNotFoundError, err)
	}
}

func (suite *UniversityServiceTestSuite) TestGetUniversityByIdAndCountryCodeGivenUniversityExists() {
	expectedUniversity := University{Id: UNIVERSITY_ID}
	suite.universityRepository.On("GetById", UNIVERSITY_ID, VALID_COUNTRY_CODE).Return(expectedUniversity, nil)

	obtainedUniversity, err := suite.service.getUniversityByIdAndCountryCode(UNIVERSITY_ID, VALID_COUNTRY_CODE)

	suite.NoError(err)
	suite.Equal(expectedUniversity, obtainedUniversity)
}

func (suite *UniversityServiceTestSuite) TestGetAllUniversities() {
	expectedUniversities := []University{*new(University)}
	suite.universityRepository.On("GetAll", VALID_COUNTRY_CODE).Return(expectedUniversities, nil)

	universities, err := suite.service.getAllUniversitiesForCountryCode(VALID_COUNTRY_CODE)

	suite.NoError(err)
	suite.Equal(expectedUniversities, universities)
}

func TestUniversityService(t *testing.T) {
	suite.Run(t, new(UniversityServiceTestSuite))
}
