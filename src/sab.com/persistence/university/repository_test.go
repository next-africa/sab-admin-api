package university

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/suite"
	"google.golang.org/appengine/datastore"
	"sab.com/domain/university"
	"sab.com/persistence"
	"testing"
	"time"
)

const (
	COUNTRY_CODE         = "CA"
	INVALID_COUNTRY_CODE = "US"
)

type UniversityRepositoryTestSuite struct {
	suite.Suite
	contextStore *persistence.ContextStore
	done         func()
	repository   universityRepository
}

func (suite *UniversityRepositoryTestSuite) SetupSuite() {
	var ctx context.Context
	ctx, suite.done = persistence.GetNewContext(suite.T())
	suite.contextStore = new(persistence.ContextStore)
	suite.contextStore.SetContext(ctx)
}

func (suite *UniversityRepositoryTestSuite) TearDownSuite() {
	suite.done()
}

func (suite *UniversityRepositoryTestSuite) SetupTest() {
	suite.repository = NewUniversityRepository(suite.contextStore)
}

func (suite *UniversityRepositoryTestSuite) TearDownTest() {
	if keys, err := datastore.NewQuery(UNIVERSITY_KIND).KeysOnly().GetAll(suite.contextStore.GetContext(), nil); err == nil {
		if len(keys) > 0 {
			datastore.DeleteMulti(suite.contextStore.GetContext(), keys)
			time.Sleep(100 * time.Millisecond)
		}
	} else {
		suite.T().Fatal(err)
	}
}

func (suite *UniversityRepositoryTestSuite) TestSaveUniversity() {
	aUniversity := createUniversity("Laval University")

	err := suite.repository.Save(&aUniversity, COUNTRY_CODE)

	suite.NoError(err)
	suite.NotNil(aUniversity.Id)
}

func (suite *UniversityRepositoryTestSuite) TestGetByIdGivenUniversityWithIdDoesNotExist() {
	_, err := suite.repository.GetById(123456789, COUNTRY_CODE)

	if suite.Error(err) {
		suite.Equal(university.UniversityNotFoundError, err)
	}
}

func (suite *UniversityRepositoryTestSuite) TestGetByIdGivenUniversityWithIdExists() {
	expectedUniversity := createUniversity("Laval University")
	suite.repository.Save(&expectedUniversity, COUNTRY_CODE)

	retrievedUniversity, err := suite.repository.GetById(expectedUniversity.Id, COUNTRY_CODE)
	suite.NoError(err)
	suite.Equal(expectedUniversity, retrievedUniversity)
}

func (suite *UniversityRepositoryTestSuite) TestGetByIdGivenUniversityWithIdExistsButCountryCodeDoesNot() {
	expectedUniversity := createUniversity("Laval University")
	suite.repository.Save(&expectedUniversity, COUNTRY_CODE)

	_, err := suite.repository.GetById(expectedUniversity.Id, INVALID_COUNTRY_CODE)

	if suite.Error(err) {
		suite.Equal(university.UniversityNotFoundError, err)
	}
}

func (suite *UniversityRepositoryTestSuite) TestGetAllUniversitiesGivenThereAreNoUniversities() {
	universities, err := suite.repository.GetAll(COUNTRY_CODE)

	suite.NoError(err)
	suite.Empty(universities)
}

func (suite *UniversityRepositoryTestSuite) TestGetAllUniversitiesGivenThereAreUniversities() {
	university1 := createUniversity("Laval University")
	university2 := createUniversity("Montreal University")
	suite.repository.Save(&university1, COUNTRY_CODE)
	suite.repository.Save(&university2, COUNTRY_CODE)

	universities, err := suite.repository.GetAll(COUNTRY_CODE)

	suite.NoError(err)
	suite.Contains(universities, university1, university2)
}

func (suite *UniversityRepositoryTestSuite) TestHasUniversityGiven() {
	university1 := createUniversity("Laval University")
	suite.repository.Save(&university1, COUNTRY_CODE)
	var fakeId int64 = 4

	time.Sleep(100 * time.Millisecond)

	testCases := []struct {
		universityId int64
		countryCode  string
		expected     bool
		msg          string
	}{
		{universityId: university1.Id, countryCode: COUNTRY_CODE, expected: true, msg: fmt.Sprintf("HasUniversity with id %d and country code %s should return true", university1.Id, COUNTRY_CODE)},
		{countryCode: COUNTRY_CODE, expected: false, msg: "HasUniversity with null id  should return false"},
		{universityId: fakeId, countryCode: COUNTRY_CODE, expected: false, msg: "HasUniversity with non existent id should return false"},
	}

	for _, testCase := range testCases {
		result, err := suite.repository.HasUniversity(testCase.universityId, testCase.countryCode)
		suite.NoError(err)
		suite.Equal(testCase.expected, result, testCase.msg)
	}
}

func createUniversity(name string) university.University {
	return university.University{
		Name:            name,
		Languages:       []string{"en", "fr"},
		Website:         "https://www.ulaval.ca",
		ProgramListLink: "https://www.ulaval.ca/les-etudes.html",
		Address: university.Address{
			Line:       "2255, Rue de l'université",
			City:       "Québec",
			State:      "QUEBEC",
			PostalCode: "G1V0A7",
		},
		Tuition: university.Tuition{
			Link:   "https://www.ulaval.ca/futurs-etudiants/couts-et-financement-des-etudes.html",
			Amount: 4000,
		},
	}
}

func TestCountryRepository(t *testing.T) {
	suite.Run(t, new(UniversityRepositoryTestSuite))
}
