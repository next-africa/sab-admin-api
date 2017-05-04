package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"golang.org/x/net/context"
	"sab.com/domain/country"
)

type CountrySchema struct {
	countryService  *country.CountryService
	nodeDefinitions *relay.NodeDefinitions

	//dynamic types
	countryType *graphql.Object

	//queries
	countriesQuery *graphql.Field
	countryQuery   *graphql.Field

	//mutations
	createCountryMutation *graphql.Field
}

func NewCountrySchema(countryService *country.CountryService, definitions *relay.NodeDefinitions) *CountrySchema {

	return &CountrySchema{countryService: countryService, nodeDefinitions: definitions}
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
					Type: graphql.NewNonNull(schema.GetCountryType()),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						if payload, ok := p.Source.(map[string]interface{}); ok {
							createdCountry := payload["country"].(country.Country)
							return newCountryNodeFromCountry(&createdCountry), nil
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
				if err := schema.saveCountry(&aCountry); err != nil {
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
				return schema.getAllCountries()
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
			Type:        graphql.NewNonNull(schema.GetCountryType()),
			Args: graphql.FieldConfigArgument{
				"code": &graphql.ArgumentConfig{
					Type:        graphql.NewNonNull(graphql.String),
					Description: "The code of the country to get",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return schema.getCountryNodeByCode(p.Args["code"].(string))
			},
		}
	}
	return schema.countryQuery
}

func (schema *CountrySchema) GetCountryByGlobalId(globalId string) (interface{}, error) {
	return schema.getCountryNodeByCode(globalId)

}

func (schema *CountrySchema) saveCountry(country *country.Country) error {
	return schema.countryService.SaveCountry(country)
}

func (schema *CountrySchema) getAllCountries() (interface{}, error) {
	if countries, err := schema.countryService.GetAllCountries(); err != nil {
		return nil, err
	} else {
		return mapCountriesToCountryNodes(countries), nil
	}
}

func (schema *CountrySchema) getCountryNodeByCode(code string) (interface{}, error) {
	theCountry, err := schema.countryService.GetCountryByCode(code)
	if err != nil {
		return nil, err
	}
	return newCountryNodeFromCountry(&theCountry), nil
}

func newCountryNodeFromCountry(aCountry *country.Country) CountryNode {
	return CountryNode{aCountry.Code, aCountry}
}

func mapCountriesToCountryNodes(countries []country.Country) []CountryNode {
	countriesMap := make([]CountryNode, len(countries))
	for i := range countries {
		countriesMap[i] = newCountryNodeFromCountry(&countries[i])
	}
	return countriesMap
}
