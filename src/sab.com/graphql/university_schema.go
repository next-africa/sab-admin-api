package graphql

import (
	"errors"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"golang.org/x/net/context"
	"sab.com/domain/university"
	"strconv"
	"strings"
)

type UniversitySchema struct {
	universityService *university.UniversityService
	nodeDefinitions   *relay.NodeDefinitions

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
	return &UniversitySchema{universityService: universityService, nodeDefinitions: definitions}
}

func (schema *UniversitySchema) GetUniversityByGlobalId(encodedGlobalId string) (interface{}, error) {
	idParts := strings.Split(encodedGlobalId, "-")

	if len(idParts) != 2 {
		return UniversityNode{}, errors.New("Invalid global university Id, the country relay Id should be of the form University:{countryCode}:{universityId}")
	}

	countryCode := idParts[0]
	universityId := idParts[1]

	return schema.getUniversityNodeByCountryCodeAndUniversityId(universityId, countryCode)
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
			Type:        graphql.NewNonNull(graphql.NewList(schema.GetUniversityType())),
			Args: graphql.FieldConfigArgument{
				"countryCode": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "The country code of the country from witch to retrieve universities",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return schema.getAllUniversities(p.Args["countryCode"].(string))
			},
		}
	}
	return schema.universitiesQuery
}

func (schema *UniversitySchema) GetUniversityQuery() *graphql.Field {
	if schema.universityQuery == nil {
		schema.universityQuery = &graphql.Field{
			Name:        "University",
			Type:        graphql.NewNonNull(schema.GetUniversityType()),
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
				return schema.getUniversityNodeByCountryCodeAndUniversityId(p.Args["universityId"].(string), p.Args["countryCode"].(string))
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
							return newUniversityNodeFromUniversity(&createdUniversity, countryCode), nil
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

				if err := schema.universityService.CreateUniversity(&aUniversity, countryCode); err != nil {
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
							return newUniversityNodeFromUniversity(&updatedUniversity, countryCode), nil
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

				if err := schema.universityService.UpdateUniversity(&aUniversity, countryCode); err != nil {
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

func (schema *UniversitySchema) getUniversityNodeByCountryCodeAndUniversityId(universityIdString string, countryCode string) (interface{}, error) {
	universityId, err := strconv.ParseInt(universityIdString, 10, 64)
	if err != nil {
		return nil, errors.New("Invalid university Id, university Id should be an Integer")
	}

	theUniversity, err := schema.universityService.GetUniversityByIdAndCountryCode(universityId, countryCode)

	if err != nil {
		return nil, err
	}

	return newUniversityNodeFromUniversity(&theUniversity, countryCode), nil
}

func (schema *UniversitySchema) getAllUniversities(countryCode string) ([]UniversityNode, error) {
	if universities, err := schema.universityService.GetAllUniversitiesForCountryCode(countryCode); err != nil {
		return []UniversityNode{}, err
	} else {
		return mapUniversitiesToUniversityNodes(universities, countryCode), nil
	}
}

func newUniversityNodeFromUniversity(aUniversity *university.University, countryCode string) UniversityNode {
	return UniversityNode{computeUniversityGlobalId(countryCode, aUniversity.Id), aUniversity}
}

func computeUniversityGlobalId(countryCode string, universityId int64) (globalId string) {
	globalId = fmt.Sprintf("%s-%v", countryCode, universityId)
	return
}

func mapUniversitiesToUniversityNodes(universities []university.University, countryCode string) []UniversityNode {
	universitiesMap := make([]UniversityNode, len(universities))
	for i := range universities {
		universitiesMap[i] = newUniversityNodeFromUniversity(&universities[i], countryCode)
	}
	return universitiesMap
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
