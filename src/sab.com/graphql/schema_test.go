package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/suite"
	"reflect"
	"sab.com/domain/country"
	"sab.com/domain/university"
	"sab.com/graphql/util"
	"testing"
)

type SchemaTestSuite struct {
	suite.Suite
	schema         *graphql.Schema
	countryRepo    *InMemoryCountryRepository
	universityRepo *InMemoryUniversityRepository
}
type T struct {
	Setup    func(suite *SchemaTestSuite)
	TearDown func()
	Name     string
	Query    string
	Schema   graphql.Schema
	Expected interface{}
}

func (suite *SchemaTestSuite) SetupTest() {
	suite.countryRepo = new(InMemoryCountryRepository)
	suite.universityRepo = new(InMemoryUniversityRepository)

	universityService := university.NewUniversityService(suite.universityRepo, suite.countryRepo)
	countryService := country.NewCountryService(suite.countryRepo)

	suite.schema = getSabGraphqlSchema(&countryService, &universityService)
}

func (suite *SchemaTestSuite) TestCountries() {

	tests := []T{
		{
			Name: "Query countries given there are no countries",
			Query: `query allCountries {
				    countries {
					id
				    }
				}
			`,
			Schema: *suite.schema,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"countries": []interface{}{},
				},
			},
		},
		{Setup: func(suite *SchemaTestSuite) {
			aCountry := country.Country{Code: "ca", Name: "Canada"}
			suite.countryRepo.Save(&aCountry)
		},
			Name: "Query countries given one country exists",
			Query: `{
				    countries {
					id
					properties {
					    code
					    name
					}
				    }
				}
			`,
			Schema: *suite.schema,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"countries": []interface{}{
						map[string]interface{}{
							"id": util.ComputeBase64("Country:"),
							"properties": map[string]interface{}{
								"code": nil,
								"name": nil,
							},
						},
						map[string]interface{}{
							"id": util.ComputeBase64("Country:ca"),
							"properties": map[string]interface{}{
								"code": "ca",
								"name": "Canada",
							},
						},
					},
				},
			},
		},
		{Setup: func(suite *SchemaTestSuite) {
			aCountry := country.Country{Code: "us", Name: "United States"}
			suite.countryRepo.Save(&aCountry)
		},
			Name: "Query countries given multiple countries exist",
			Query: `{
				    countries {
					id
					properties {
					    code
					    name
					}
				    }
				}
			`,
			Schema: *suite.schema,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"countries": []interface{}{
						map[string]interface{}{
							"id": util.ComputeBase64("Country:"),
							"properties": map[string]interface{}{
								"code": nil,
								"name": nil,
							},
						},
						map[string]interface{}{
							"id": util.ComputeBase64("Country:"),
							"properties": map[string]interface{}{
								"code": nil,
								"name": nil,
							},
						},
						map[string]interface{}{
							"id": util.ComputeBase64("Country:ca"),
							"properties": map[string]interface{}{
								"code": "ca",
								"name": "Canada",
							},
						},
						map[string]interface{}{
							"id": util.ComputeBase64("Country:us"),
							"properties": map[string]interface{}{
								"code": "us",
								"name": "United States",
							},
						},
					},
				},
			},
		},
	}

	for _, t := range tests {
		suite.testGraphql(t)
	}
}

func (suite *SchemaTestSuite) testGraphql(test T) {
	defer func() {
		if test.TearDown != nil {
			test.TearDown()
		}
	}()
	if test.Setup != nil {
		test.Setup(suite)
	}

	params := graphql.Params{
		Schema:        test.Schema,
		RequestString: test.Query,
	}

	result := graphql.Do(params)
	if len(result.Errors) > 0 {
		suite.T().Fatalf("%v, wrong result, unexpected errors: %v", test.Name, result.Errors)
	}

	if !reflect.DeepEqual(result, test.Expected) {
		suite.T().Fatalf("%v, wrong result, query: %v, graphql result diff: %v", test.Name, test.Query, Diff(test.Expected, result))
	}
}

func TestSchema(t *testing.T) {
	suite.Run(t, new(SchemaTestSuite))
}
