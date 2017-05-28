package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/suite"
	"reflect"
	"sab.com/domain/country"
	"sab.com/domain/university"
)

type TestCase struct {
	Setup                   func()
	TearDown                func()
	Name                    string
	Query                   string
	VariableValues          map[string]interface{}
	Expected                *graphql.Result
	ValidateMutationSuccess func() error
}

type TestMutation struct {
	Setup    func()
	TearDown func()
	Name     string
	Mutation string
	Expected interface{}
}

type SchemaTestSuite struct {
	suite.Suite
	schema            *graphql.Schema
	countryService    country.CountryService
	universityService university.UniversityService
}

func (s *SchemaTestSuite) testGraphql(test TestCase) {
	defer func() {
		if test.TearDown != nil {
			test.TearDown()
		}
	}()

	if test.Setup != nil {
		test.Setup()
	}

	params := graphql.Params{
		Schema:         *s.schema,
		RequestString:  test.Query,
		VariableValues: test.VariableValues,
	}

	result := graphql.Do(params)

	if test.Expected.Errors == nil && result.Errors != nil {
		s.T().Fatalf("%v, query: %v, got unexpected errors: %v", test.Name, test.Query, result.Errors)
	}

	if !reflect.DeepEqual(result, test.Expected) {
		s.T().Fatalf("%v, wrong result, query: %v, graphql result diff: %v", test.Name, test.Query, Diff(test.Expected, result))
	}

	if test.ValidateMutationSuccess != nil {
		if err := test.ValidateMutationSuccess(); err != nil {
			s.T().Fatalf("%v, query: %v, Mutation failed with erros: %v", test.Name, test.Query, err)
		}
	}
}

func (s *SchemaTestSuite) SetupTest() {
	inMemoryUniversityRepository := new(InMemoryUniversityRepository)
	inMemoryCountryRepository := new(InMemoryCountryRepository)
	s.universityService = university.NewUniversityService(inMemoryUniversityRepository, inMemoryCountryRepository)
	s.countryService = country.NewCountryService(inMemoryCountryRepository)

	s.schema = getSabGraphqlSchema(&s.countryService, &s.universityService)
}
