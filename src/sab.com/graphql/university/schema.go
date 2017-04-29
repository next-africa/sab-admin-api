package graphql

import (
	"errors"
	"github.com/next-africa/graphql-go"
	"github.com/next-africa/graphql-go-relay"
	"golang.org/x/net/context"
	"sab.com/domain/university"
	"strconv"
)

type UniversitySchema struct {
	universityGraphService *UniversityGraphService
	nodeDefinitions        *relay.NodeDefinitions

	//dynamic types
	universityType *graphql.Object

	//queries
	universitiesQuery *graphql.Field
	universityQuery   *graphql.Field

	//mutations
	createUniversityMutation *graphql.Field
	updateUniversityMutation *graphql.Field
}

func NewUniversitySchema(universityService *university.UniversityService, definitions *relay.NodeDefinitions) *UniversitySchema {
	universityGraphService := NewUniversityGraphqlService(universityService)
	return &UniversitySchema{universityGraphService: &universityGraphService, nodeDefinitions: definitions}
}

func (schema *UniversitySchema) GetUniversityByGlobalId(encodedGlobalId string) (UniversityNode, error) {
	return schema.universityGraphService.GetUniversityByGlobalId(encodedGlobalId)
}

func (schema *UniversitySchema) GetUniversityType() *graphql.Object {
	if schema.universityType == nil {
		schema.universityType = graphql.NewObject(graphql.ObjectConfig{
			Name:        "University",
			Description: "A University",
			Fields: graphql.Fields{
				"id": relay.GlobalIDField("University", nil),
				"properties": &graphql.Field{
					Type:        UniversityPropertiesType,
					Description: "The properties of the university",
				},
			},
			Interfaces: []*graphql.Interface{
				schema.nodeDefinitions.NodeInterface,
			},
		})
	}
	return schema.universityType
}

func (schema *UniversitySchema) GetUniversitiesQuery() *graphql.Field {
	if schema.universitiesQuery == nil {
		schema.universitiesQuery = &graphql.Field{
			Name:        "Universities",
			Description: "Get the list of universities of this the given country",
			Type:        graphql.NewList(schema.GetUniversityType()),
			Args: graphql.FieldConfigArgument{
				"countryCode": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "The country code of the country from witch to retrieve universities",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return schema.universityGraphService.GetAllUniversities(p.Args["countryCode"].(string))
			},
		}
	}
	return schema.universitiesQuery
}

func (schema *UniversitySchema) GetUniversityQuery() *graphql.Field {
	if schema.universityQuery == nil {
		schema.universityQuery = &graphql.Field{
			Name:        "University",
			Type:        schema.GetUniversityType(),
			Description: "Get a country by country code",
			Args: graphql.FieldConfigArgument{
				"countryCode": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "The code of the country to get",
				},
				"universityId": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "The university Id to get",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return schema.universityGraphService.GetUniversityNodeByCountryCodeAndUniversityId(p.Args["universityId"].(string), p.Args["countryCode"].(string))
			},
		}
	}
	return schema.universityQuery
}

func (schema *UniversitySchema) GetCreateUniversityMutation() *graphql.Field {
	if schema.createUniversityMutation == nil {
		schema.createUniversityMutation = relay.MutationWithClientMutationID(relay.MutationConfig{
			Name:        "CreateUniversity",
			InputFields: CreateUniversityInputFields,
			OutputFields: graphql.Fields{
				"university": &graphql.Field{
					Type: schema.GetUniversityType(),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						if payload, ok := p.Source.(map[string]interface{}); ok {
							createdUniversity := payload["university"].(university.University)
							countryCode := payload["countryCode"].(string)
							return schema.universityGraphService.NewUniversityNodeFromUniversity(&createdUniversity, countryCode), nil
						} else {
							return nil, nil
						}
					},
				},
			},

			MutateAndGetPayload: func(inputMap map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
				countryCode := inputMap["countryCode"].(string)

				aUniversity, err := getUniversityFromInputMap(inputMap["university"].(map[string]interface{}))

				if err != nil {
					return nil, err
				}

				if err := schema.universityGraphService.CreateUniversity(&aUniversity, countryCode); err != nil {
					return nil, err
				}

				return map[string]interface{}{
					"university":  aUniversity,
					"countryCode": countryCode,
				}, nil
			},
		})
	}

	return schema.createUniversityMutation
}

func getUniversityFromInputMap(inputMap map[string]interface{}) (university.University, error) {
	name := inputMap["name"].(string)
	website := inputMap["website"].(string)

	languageArray := inputMap["languages"].([]interface{})
	languages := make([]string, len(languageArray))
	for i, v := range languageArray {
		languages[i] = v.(string)
	}

	var programListLink string
	if programListLinkInput, ok := inputMap["programListLink"]; ok {
		programListLink = programListLinkInput.(string)
	}

	var tuition university.Tuition
	if tuitionInput, ok := inputMap["tuition"]; ok {
		tuitionMap := tuitionInput.(map[string]interface{})
		tuition = university.Tuition{
			Link:   tuitionMap["link"].(string),
			Amount: tuitionMap["amount"].(int),
		}
	}

	var address university.Address
	if addressInput, ok := inputMap["address"]; ok {
		addressMap := addressInput.(map[string]interface{})

		var line string
		if lineInput, ok := addressMap["line"]; ok {
			line = lineInput.(string)
		}

		var city string
		if cityInput, ok := addressMap["city"]; ok {
			city = cityInput.(string)
		}

		var state string
		if stateInput, ok := addressMap["state"]; ok {
			state = stateInput.(string)
		}

		var postalCode string
		if postalCodeInput, ok := addressMap["postalCode"]; ok {
			postalCode = postalCodeInput.(string)
		}

		address = university.Address{
			Line:       line,
			City:       city,
			State:      state,
			PostalCode: postalCode,
		}
	}

	var universityId int64

	if idInput, ok := inputMap["id"]; ok {
		var err error
		universityId, err = strconv.ParseInt(idInput.(string), 10, 64)
		if err != nil {
			return university.University{}, errors.New("The university id input is not valid")
		}
	}

	return university.University{
		Id:              universityId,
		Name:            name,
		Languages:       languages,
		Website:         website,
		ProgramListLink: programListLink,
		Address:         address,
		Tuition:         tuition,
	}, nil
}

func (schema *UniversitySchema) GetUpdateUniversityMutation() *graphql.Field {
	if schema.updateUniversityMutation == nil {
		schema.updateUniversityMutation = relay.MutationWithClientMutationID(relay.MutationConfig{
			Name:        "UpdateUniversity",
			InputFields: UpdateUniversityInputFields,
			OutputFields: graphql.Fields{
				"university": &graphql.Field{
					Type: schema.GetUniversityType(),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						if payload, ok := p.Source.(map[string]interface{}); ok {
							updatedUniversity := payload["university"].(university.University)
							countryCode := payload["countryCode"].(string)
							return schema.universityGraphService.NewUniversityNodeFromUniversity(&updatedUniversity, countryCode), nil
						} else {
							return nil, nil
						}
					},
				},
			},

			MutateAndGetPayload: func(inputMap map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
				countryCode := inputMap["countryCode"].(string)

				aUniversity, err := getUniversityFromInputMap(inputMap["university"].(map[string]interface{}))

				if err != nil {
					return nil, err
				}

				if err := schema.universityGraphService.UpdateUniversity(&aUniversity, countryCode); err != nil {
					return nil, err
				}

				return map[string]interface{}{
					"university":  aUniversity,
					"countryCode": countryCode,
				}, nil
			},
		})
	}

	return schema.updateUniversityMutation
}
