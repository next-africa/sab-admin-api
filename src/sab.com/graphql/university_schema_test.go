package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/stretchr/testify/suite"
	"sab.com/domain/country"
	"sab.com/domain/university"
	"testing"
)

type UniversitySchemaTestSuite struct {
	SchemaTestSuite
}

func createDummyUniversityObject(name string) university.University {
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

func (s *UniversitySchemaTestSuite) TestUniversitiesQuery() {

	tests := []TestCase{
		{
			Name: "Query universities given an invalid country code",
			Query: `query allUniversities {
				    universities(countryCode: "ca") {
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
			Name: "Query universities given a valid country code but there are no universities for given country code",
			Setup: func() {
				existingCountry := country.Country{Name: "Canada", Code: "ca"}
				s.countryService.SaveCountry(&existingCountry)
			},
			Query: `query allUniversities {
				    universities(countryCode: "ca") {
					id
				    }
				}
			`,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"universities": []interface{}{},
				},
			},
		},
		{
			Setup: func() {
				aUniversity := createDummyUniversityObject("Laval University")
				s.universityService.CreateUniversity(&aUniversity, "ca")
			},
			Name: "Query universities given one university exists",
			Query: `{
				    universities(countryCode: "ca") {
					id
					properties {
					    id
					    name
					    languages
					    website
					    programListLink
					    address {
					        line
					        city
					        state
					        postalCode
			                    }
			                    tuition {
			                        amount
			                        link
			                    }
					}
				    }
				}
			`,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"universities": []interface{}{
						map[string]interface{}{
							"id": ComputeBase64("University:ca-1"),
							"properties": map[string]interface{}{
								"id":              "1",
								"name":            "Laval University",
								"languages":       []interface{}{"en", "fr"},
								"website":         "https://www.ulaval.ca",
								"programListLink": "https://www.ulaval.ca/les-etudes.html",
								"address": map[string]interface{}{
									"line":       "2255, Rue de l'université",
									"city":       "Québec",
									"state":      "QUEBEC",
									"postalCode": "G1V0A7",
								},
								"tuition": map[string]interface{}{
									"link":   "https://www.ulaval.ca/futurs-etudiants/couts-et-financement-des-etudes.html",
									"amount": 4000,
								},
							},
						},
					},
				},
			},
		},
		{
			Setup: func() {
				secondUniversity := createDummyUniversityObject("Montreal University")
				s.universityService.CreateUniversity(&secondUniversity, "ca")
			},
			Name: "Query universities given multiple universties exist",
			Query: `{
				    universities(countryCode: "ca") {
					id
					properties {
					    name
					}
				    }
				}
			`,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"universities": []interface{}{
						map[string]interface{}{
							"id": ComputeBase64("University:ca-1"),
							"properties": map[string]interface{}{
								"name": "Laval University",
							},
						},
						map[string]interface{}{
							"id": ComputeBase64("University:ca-2"),
							"properties": map[string]interface{}{
								"name": "Montreal University",
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

func (s *UniversitySchemaTestSuite) TestUniversityQuery() {
	tests := []TestCase{
		{
			Name: "Query university given a non existent country code",
			Query: `{
				    university(countryCode: "ca", universityId: "123456789") {
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
			Name: "Query university given an existent country code but non existent id",
			Setup: func() {
				aCountry := country.Country{Code: "ca", Name: "Canada"}
				s.countryService.SaveCountry(&aCountry)
			},
			Query: `{
				    university(countryCode: "ca", universityId: "123456789") {
					id
				    }
				}
			`,
			Expected: &graphql.Result{
				Data: nil,
				Errors: []gqlerrors.FormattedError{
					gqlerrors.NewFormattedError(university.UniversityNotFoundError.Error()),
				},
			},
		},
		{
			Name: "Query university given both existent country code and existent id",
			Setup: func() {
				aUniversity := createDummyUniversityObject("Laval University")
				s.universityService.CreateUniversity(&aUniversity, "ca")
			},
			Query: `{
				    university(countryCode: "ca", universityId: "1") {
					id
					properties {
					    id
					    name
					    languages
					    website
					    programListLink
					    address {
					        line
					        city
					        state
					        postalCode
			                    }
			                    tuition {
			                        amount
			                        link
			                    }
					}
				    }
				}
			`,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"university": map[string]interface{}{
						"id": ComputeBase64("University:ca-1"),
						"properties": map[string]interface{}{
							"id":              "1",
							"name":            "Laval University",
							"languages":       []interface{}{"en", "fr"},
							"website":         "https://www.ulaval.ca",
							"programListLink": "https://www.ulaval.ca/les-etudes.html",
							"address": map[string]interface{}{
								"line":       "2255, Rue de l'université",
								"city":       "Québec",
								"state":      "QUEBEC",
								"postalCode": "G1V0A7",
							},
							"tuition": map[string]interface{}{
								"link":   "https://www.ulaval.ca/futurs-etudiants/couts-et-financement-des-etudes.html",
								"amount": 4000,
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

func (s *UniversitySchemaTestSuite) TestUniversityNodeQuery() {
	tests := []TestCase{
		{
			Name: "Querying university node given a non existent country code",
			Query: `query University($id: ID!){
					node(id: $id){
						id
				 	}
				}
			`,
			VariableValues: map[string]interface{}{
				"id": ComputeBase64("University:ca-1"),
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
			Name: "Querying university node given existent country code but non existent university id",
			Setup: func() {
				aCountry := country.Country{Code: "ca", Name: "Canada"}
				s.countryService.SaveCountry(&aCountry)
			},
			Query: `query University($id: ID!){
					node(id: $id){
						id
				 	}
				}
			`,
			VariableValues: map[string]interface{}{
				"id": ComputeBase64("University:ca-1"),
			},
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"node": nil,
				},
				Errors: []gqlerrors.FormattedError{
					gqlerrors.NewFormattedError(university.UniversityNotFoundError.Error()),
				},
			},
		},
		{
			Name: "Querying a university node given both existent country code and university id ",
			Setup: func() {
				aUniversity := createDummyUniversityObject("Laval University")
				s.universityService.CreateUniversity(&aUniversity, "ca")
			},
			Query: `query Country($id: ID!){
					node(id: $id) {
						id
						... on University {
							properties {
								name
							}
						}
					}
				}
			`,
			VariableValues: map[string]interface{}{
				"id": ComputeBase64("University:ca-1"),
			},
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"node": map[string]interface{}{
						"id": ComputeBase64("University:ca-1"),
						"properties": map[string]interface{}{
							"name": "Laval University",
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

func (s *UniversitySchemaTestSuite) TestCreateUniversityMutation() {
	tests := []TestCase{
		{
			Name: "Creating a University",
			Setup: func() {
				existingCountry := country.Country{Name: "Canada", Code: "ca"}
				s.countryService.SaveCountry(&existingCountry)
			},
			Query: `mutation CreateUniversity ($universityInput: CreateUniversityInput!){
					createUniversity(input: $universityInput) {
						university {
							id
							properties {
							    id
							    name
							    languages
							    website
							    programListLink
							    address {
								line
								city
								state
								postalCode
							    }
							    tuition {
								amount
								link
							    }
							}
						}
					}
				}
			`,
			VariableValues: map[string]interface{}{
				"universityInput": map[string]interface{}{
					"clientMutationId": "abcxyz",
					"countryCode":      "ca",
					"university": map[string]interface{}{
						"name":            "Laval University",
						"languages":       []string{"en", "fr"},
						"website":         "https://www.ulaval.ca",
						"programListLink": "https://www.ulaval.ca/les-etudes.html",
						"address": map[string]interface{}{
							"line":       "2255, Rue de l'université",
							"city":       "Québec",
							"state":      "QUEBEC",
							"postalCode": "G1V0A7",
						},
						"tuition": map[string]interface{}{
							"amount": 4000,
							"link":   "https://www.ulaval.ca/futurs-etudiants/couts-et-financement-des-etudes.html",
						},
					},
				},
			},
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"createUniversity": map[string]interface{}{
						"university": map[string]interface{}{
							"id": ComputeBase64("University:ca-1"),
							"properties": map[string]interface{}{
								"id":              "1",
								"name":            "Laval University",
								"languages":       []interface{}{"en", "fr"},
								"website":         "https://www.ulaval.ca",
								"programListLink": "https://www.ulaval.ca/les-etudes.html",
								"address": map[string]interface{}{
									"line":       "2255, Rue de l'université",
									"city":       "Québec",
									"state":      "QUEBEC",
									"postalCode": "G1V0A7",
								},
								"tuition": map[string]interface{}{
									"link":   "https://www.ulaval.ca/futurs-etudiants/couts-et-financement-des-etudes.html",
									"amount": 4000,
								},
							},
						},
					},
				},
			},
			ValidateMutationSuccess: func() error {
				_, err := s.universityService.GetUniversityByIdAndCountryCode(1, "ca")
				return err
			},
		},
	}

	for _, t := range tests {
		s.testGraphql(t)
	}
}

func (s *UniversitySchemaTestSuite) TestUpdateUniversityMutation() {
	tests := []TestCase{
		{
			Name: "Updating a University",
			Setup: func() {
				existingCountry := country.Country{Name: "Canada", Code: "ca"}
				s.countryService.SaveCountry(&existingCountry)

				aUniversity := createDummyUniversityObject("Laval University")
				s.universityService.CreateUniversity(&aUniversity, "ca")
			},
			Query: `mutation UpdateUniversity ($universityInput: UpdateUniversityInput!){
					updateUniversity(input: $universityInput) {
						university {
							id
							properties {
							    id
							    name
							    languages
							    website
							}
						}
					}
				}
			`,
			VariableValues: map[string]interface{}{
				"universityInput": map[string]interface{}{
					"clientMutationId": "abcxyz",
					"countryCode":      "ca",
					"university": map[string]interface{}{
						"id":        "1",
						"name":      "Montreal University",
						"languages": []string{"de", "fr"},
						"website":   "https://www.ulaval.ca",
					},
				},
			},
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"updateUniversity": map[string]interface{}{
						"university": map[string]interface{}{
							"id": ComputeBase64("University:ca-1"),
							"properties": map[string]interface{}{
								"id":        "1",
								"name":      "Montreal University",
								"languages": []interface{}{"de", "fr"},
								"website":   "https://www.ulaval.ca",
							},
						},
					},
				},
			},
			ValidateMutationSuccess: func() error {
				_, err := s.universityService.GetUniversityByIdAndCountryCode(1, "ca")
				return err
			},
		},
	}

	for _, t := range tests {
		s.testGraphql(t)
	}
}

func TestUniversitySchema(t *testing.T) {
	suite.Run(t, new(UniversitySchemaTestSuite))
}
