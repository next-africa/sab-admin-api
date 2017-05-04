package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/stretchr/testify/suite"
	"sab.com/domain/country"
	"testing"
)

type CountrySchemaTestSuite struct {
	SchemaTestSuite
}

func (s *CountrySchemaTestSuite) TestCountriesQuery() {

	tests := []TestCase{
		{
			Name: "Query countries given there are no countries",
			Query: `query allCountries {
				    countries {
					id
				    }
				}
			`,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"countries": []interface{}{},
				},
			},
		},
		{
			Setup: func() {
				aCountry := country.Country{Code: "ca", Name: "Canada"}
				s.countryService.SaveCountry(&aCountry)
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
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"countries": []interface{}{
						map[string]interface{}{
							"id": ComputeBase64("Country:ca"),
							"properties": map[string]interface{}{
								"code": "ca",
								"name": "Canada",
							},
						},
					},
				},
			},
		},
		{
			Setup: func() {
				aCountry := country.Country{Code: "us", Name: "United States"}
				s.countryService.SaveCountry(&aCountry)
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
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"countries": []interface{}{
						map[string]interface{}{
							"id": ComputeBase64("Country:ca"),
							"properties": map[string]interface{}{
								"code": "ca",
								"name": "Canada",
							},
						},
						map[string]interface{}{
							"id": ComputeBase64("Country:us"),
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
		s.testGraphql(t)
	}
}

func (s *CountrySchemaTestSuite) TestCountryQuery() {
	tests := []TestCase{
		{
			Name: "Querying a country given invalid code",
			Query: `{
				    country(code: "ca"){
				    	id
				    }
				}
			`,
			Expected: &graphql.Result{
				Data: nil,
				Errors: []gqlerrors.FormattedError{
					gqlerrors.NewFormattedError(country.CountryNotFoundError.Error()),
				},
			},
		},

		{
			Setup: func() {
				aCountry := country.Country{Code: "ca", Name: "Canada"}
				s.countryService.SaveCountry(&aCountry)
			},
			Name: "Querying a country given valid code",
			Query: `{
					country(code: "ca") {
						id
						properties {
							code
							name
						}
					}
				}
			`,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"country": map[string]interface{}{
						"id": ComputeBase64("Country:ca"),
						"properties": map[string]interface{}{
							"code": "ca",
							"name": "Canada",
						},
					},
				},
			},
		},
	}

	for _, t := range tests {
		s.testGraphql(t)
	}
}

func (s *CountrySchemaTestSuite) TestCountryNodeQuery() {
	tests := []TestCase{
		{
			Name: "Querying country node with invalid code",
			Query: `query Country($id: ID!){
					node(id: $id){
						id
				 	}
				}
			`,
			VariableValues: map[string]interface{}{
				"id": ComputeBase64("Country:ca"),
			},
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"node": nil,
				},
				Errors: []gqlerrors.FormattedError{
					gqlerrors.NewFormattedError(country.CountryNotFoundError.Error()),
				},
			},
		},
		{
			Setup: func() {
				aCountry := country.Country{Code: "ca", Name: "Canada"}
				s.countryService.SaveCountry(&aCountry)
			},
			Name: "Querying a country node given valid code",
			Query: `query Country($id: ID!){
					node(id: $id) {
						id
						... on Country {
							properties {
								code
								name
							}
						}
					}
				}
			`,
			VariableValues: map[string]interface{}{
				"id": ComputeBase64("Country:ca"),
			},
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"node": map[string]interface{}{
						"id": ComputeBase64("Country:ca"),
						"properties": map[string]interface{}{
							"code": "ca",
							"name": "Canada",
						},
					},
				},
			},
		},
	}

	for _, t := range tests {
		s.testGraphql(t)
	}
}

func (s *CountrySchemaTestSuite) TestCreateCountryMutation() {
	tests := []TestCase{
		{
			Name: "Creating a Country",
			Query: `mutation CreateCountry ($countryInput: CreateCountryInput!){
					createCountry(input: $countryInput) {
						country {
							id
							properties {
								code
								name
							}
						}
					}
				}
			`,
			VariableValues: map[string]interface{}{
				"countryInput": map[string]interface{}{
					"clientMutationId": "abcxyz",
					"code":             "us",
					"name":             "United States",
				},
			},
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"createCountry": map[string]interface{}{
						"country": map[string]interface{}{
							"id": ComputeBase64("Country:us"),
							"properties": map[string]interface{}{
								"code": "us",
								"name": "United States",
							},
						},
					},
				},
			},
			ValidateMutationSuccess: func() error {
				_, err := s.countryService.GetCountryByCode("us")
				return err
			},
		},
	}

	for _, t := range tests {
		s.testGraphql(t)
	}

}

func TestSchema(t *testing.T) {
	suite.Run(t, new(CountrySchemaTestSuite))
}
