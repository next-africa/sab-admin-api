package graphql

import (
	"github.com/next-africa/graphql-go"
	"github.com/next-africa/graphql-go-relay"
	"golang.org/x/net/context"
	"sab.com/domain/country"
)

type CountrySchema struct {
	countryGraphService *CountryGraphService
	nodeDefinitions     *relay.NodeDefinitions

	//dynamic types
	countryType *graphql.Object

	//queries
	countriesQuery *graphql.Field
	countryQuery   *graphql.Field

	//mutations
	createCountryMutation *graphql.Field
}

func NewCountrySchema(countryService *country.CountryService, definitions *relay.NodeDefinitions) *CountrySchema {
	countryGraphService := NewCountryGraphqlService(countryService)
	return &CountrySchema{countryGraphService: &countryGraphService, nodeDefinitions: definitions}
}

func (schema *CountrySchema) GetCountryType() *graphql.Object {
	if schema.countryType == nil {
		schema.countryType = graphql.NewObject(graphql.ObjectConfig{
			Name:        "Country",
			Description: "A country",
			Fields: graphql.Fields{
				"id": relay.GlobalIDField("Country", nil),
				"properties": &graphql.Field{
					Type:        CountryPropertiesType,
					Description: "The properties of the country",
				},
			},
			Interfaces: []*graphql.Interface{
				schema.nodeDefinitions.NodeInterface,
			},
		})
	}

	return schema.countryType
}

func (schema *CountrySchema) GetCreateCountryMutation() *graphql.Field {
	if schema.createCountryMutation == nil {
		schema.createCountryMutation = relay.MutationWithClientMutationID(relay.MutationConfig{
			Name:        "CreateCountry",
			InputFields: CreateCountryInputFields,
			OutputFields: graphql.Fields{
				"country": &graphql.Field{
					Type: schema.GetCountryType(),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						if payload, ok := p.Source.(map[string]interface{}); ok {
							createdCountry := payload["country"].(country.Country)
							return schema.countryGraphService.NewCountryNodeFromCountry(&createdCountry), nil
						} else {
							return nil, nil
						}
					},
				},
			},
			MutateAndGetPayload: func(inputMap map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
				code := inputMap["code"].(string)
				name := inputMap["name"].(string)
				aCountry := country.Country{Code: code, Name: name}
				if err := schema.countryGraphService.SaveCountry(&aCountry); err != nil {
					return nil, err
				}
				return map[string]interface{}{
					"country": aCountry,
				}, nil
			},
		})
	}

	return schema.createCountryMutation
}

func (schema *CountrySchema) GetCountriesQuery() *graphql.Field {
	if schema.countriesQuery == nil {
		schema.countriesQuery = &graphql.Field{
			Name:        "Countries",
			Description: "Get the list of countries supported by the application",
			Type:        graphql.NewList(schema.GetCountryType()),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return schema.countryGraphService.GetAllCountries()
			},
		}
	}
	return schema.countriesQuery
}

func (schema *CountrySchema) GetCountryQuery() *graphql.Field {
	if schema.countryQuery == nil {
		schema.countryQuery = &graphql.Field{
			Name:        "Country",
			Description: "Get a country by country code",
			Type:        schema.GetCountryType(),
			Args: graphql.FieldConfigArgument{
				"code": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "The code of the country to get",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return schema.countryGraphService.GetCountryNodeByCode(p.Args["code"].(string))
			},
		}
	}
	return schema.countryQuery
}

func (schema *CountrySchema) GetCountryByGlobalId(encodedGlobalId string) (CountryNode, error) {
	return schema.countryGraphService.GetCountryByGlobalId(encodedGlobalId)
}
