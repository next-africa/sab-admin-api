package graphql

import (
	"errors"
	"github.com/next-africa/graphql-go"
	"github.com/next-africa/graphql-go-relay"
	"golang.org/x/net/context"
	"sab.com/domain/country"
	"sab.com/domain/university"
	graphqlCountry "sab.com/graphql/country"
	graphqlUniversity "sab.com/graphql/university"
)

var nodeDefinitions *relay.NodeDefinitions

// exported schema, defined in init()
var schema *graphql.Schema

var countrySchema *graphqlCountry.CountrySchema
var universitySchema *graphqlUniversity.UniversitySchema

func GetSabGraphqlSchema(countryService *country.CountryService, universityService *university.UniversityService) *graphql.Schema {
	if schema == nil {
		var err error
		schema, err = createSchema(countryService, universityService)
		if err != nil {
			panic(err)
		}
	}

	return schema
}

func createSchema(countryService *country.CountryService, universityService *university.UniversityService) (*graphql.Schema, error) {
	nodeDefinitions = relay.NewNodeDefinitions(relay.NodeDefinitionsConfig{
		IDFetcher: func(id string, info graphql.ResolveInfo, ctx context.Context) (interface{}, error) {
			// resolve id from global id
			resolvedID := relay.FromGlobalID(id)

			// based on id and its type, return the object
			switch resolvedID.Type {
			case "Country":
				return countrySchema.GetCountryByGlobalId(resolvedID.ID)
			case "University":
				return universitySchema.GetUniversityByGlobalId(resolvedID.ID)
			default:
				return nil, errors.New("Unknown node type")
			}
		},

		TypeResolve: func(p graphql.ResolveTypeParams) *graphql.Object {
			// based on the type of the value, return GraphQLObjectType
			switch p.Value.(type) {
			case graphqlUniversity.UniversityNode:
				return universitySchema.GetUniversityType()
			default:
				return countrySchema.GetCountryType()
			}
		},
	})

	countrySchema = graphqlCountry.NewCountrySchema(countryService, nodeDefinitions)
	universitySchema = graphqlUniversity.NewUniversitySchema(universityService, nodeDefinitions)

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "SabQuery",
		Fields: graphql.Fields{
			"countries":    countrySchema.GetCountriesQuery(),
			"country":      countrySchema.GetCountryQuery(),
			"universities": universitySchema.GetUniversitiesQuery(),
			"university":   universitySchema.GetUniversityQuery(),
			"node":         nodeDefinitions.NodeField,
		},
	})

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "SabMutation",
		Description: "All the mutation available on Study abroad apy",
		Fields: graphql.Fields{
			"createCountry":    countrySchema.GetCreateCountryMutation(),
			"createUniversity": universitySchema.GetCreateUniversityMutation(),
		},
	})

	aSchema, aErr := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})

	return &aSchema, aErr
}
