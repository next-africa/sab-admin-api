package country

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/suite"
	"google.golang.org/appengine/datastore"
	"sab.com/domain/country"
	"sab.com/persistence"
	"strings"
	"testing"
	"time"
)

const COUNTRY_CODE = "CA"
const COUNTRY_NAME = "CANADA"

const COUNTRY_CODE_1 = "US"
const COUNTRY_NAME_1 = "United States"
const NON_EXISTING_CODE = "MOCHO"

type CountryRepositoryTestSuite struct {
	suite.Suite
	contextStore *persistence.ContextStore
	done         func()
	repository   countryRepository
}

func (suite *CountryRepositoryTestSuite) SetupSuite() {
	var ctx context.Context
	ctx, suite.done = persistence.GetNewContext(suite.T())
	suite.contextStore = new(persistence.ContextStore)
	suite.contextStore.SetContext(ctx)
}

func (suite *CountryRepositoryTestSuite) TearDownSuite() {
	suite.done()
}

func (suite *CountryRepositoryTestSuite) SetupTest() {
	suite.repository = NewCountryRepository(suite.contextStore)
}

func (suite *CountryRepositoryTestSuite) TearDownTest() {
	if keys, err := datastore.NewQuery(COUNTRY_KIND).KeysOnly().GetAll(suite.contextStore.GetContext(), nil); err == nil {
		if len(keys) > 0 {
			datastore.DeleteMulti(suite.contextStore.GetContext(), keys)
			time.Sleep(100 * time.Millisecond)
		}
	} else {
		suite.T().Fatal(err)
	}
}

func (suite *CountryRepositoryTestSuite) TestSaveCountry() {

	aCountry := country.Country{Code: COUNTRY_CODE, Name: COUNTRY_NAME}

	err := suite.repository.Save(&aCountry)

	suite.NoError(err)

	savedCountry := country.Country{}

	key := datastore.NewKey(suite.contextStore.GetContext(), "Country", strings.ToLower(COUNTRY_CODE), 0, nil)
	err1 := datastore.Get(suite.contextStore.GetContext(), key, &savedCountry)

	suite.NoError(err1)

	suite.Equal(aCountry, savedCountry)
}

func (suite *CountryRepositoryTestSuite) TestGetByCodeGivenCountryDoesNotExist() {
	_, err := suite.repository.GetByCode(NON_EXISTING_CODE)

	if suite.Error(err) {
		suite.Equal(country.CountryNotFoundError, err)
	}
}

func (suite *CountryRepositoryTestSuite) TestGetByCodeGivenCountryExists() {
	existingCountry := country.Country{Code: COUNTRY_CODE, Name: COUNTRY_NAME}
	suite.repository.Save(&existingCountry)

	theCountry, err := suite.repository.GetByCode(COUNTRY_CODE)

	suite.NoError(err)
	suite.Equal(existingCountry, theCountry)
}

func (suite *CountryRepositoryTestSuite) TestGetAllCountriesGivenThereAreNoCountries() {
	countries, err := suite.repository.GetAll()

	suite.NoError(err)
	suite.Empty(countries)
}

func (suite *CountryRepositoryTestSuite) TestGetAllCountriesGivenThereAreCountries() {
	country1 := country.Country{Code: COUNTRY_CODE, Name: COUNTRY_NAME}
	country2 := country.Country{Code: COUNTRY_CODE_1, Name: COUNTRY_NAME_1}
	suite.repository.Save(&country1)
	suite.repository.Save(&country2)
	time.Sleep(100 * time.Millisecond)

	countries, err := suite.repository.GetAll()

	suite.NoError(err)
	suite.Contains(countries, country1, country2)
}

func (suite *CountryRepositoryTestSuite) TestHasCountryWithCode() {

	suite.repository.Save(&country.Country{Code: COUNTRY_CODE, Name: COUNTRY_NAME})

	testCases := []struct {
		code     string
		expected bool
		msg      string
	}{
		{COUNTRY_CODE, true, fmt.Sprintf("HasCountryWithCode with code %s should return true", COUNTRY_CODE)},
		{COUNTRY_CODE_1, false, fmt.Sprintf("HasCountryWithCode with code %s should return false", COUNTRY_CODE_1)},
	}

	for _, testCase := range testCases {
		result, err := suite.repository.HasCountryWithCode(testCase.code)
		suite.NoError(err)
		suite.Equal(testCase.expected, result, testCase.msg)
	}

}

func TestCountryRepository(t *testing.T) {
	suite.Run(t, new(CountryRepositoryTestSuite))
}
