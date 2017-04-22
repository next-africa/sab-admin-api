package graphql

import (
	"errors"
	"github.com/next-africa/graphql-go"
	"github.com/next-africa/graphql-go-relay"
	"golang.org/x/net/context"
	"sab.com/domain/country"
	"sab.com/domain/university"
)

var nodeDefinitions *relay.NodeDefinitions
var universityType *graphql.Object
var countryType *graphql.Object

var countryService *country.CountryService
var universityService *university.UniversityService

// exported schema, defined in init()
var schema *graphql.Schema

func GetSabGraphqlSchema(ctryService *country.CountryService, univService *university.UniversityService) *graphql.Schema {
	if schema == nil {
		countryService = ctryService
		universityService = univService
		var err error
		schema, err = createSchema()
		if err != nil {
			panic(err)
		}
	}

	return schema
}

func createSchema() (*graphql.Schema, error) {
	nodeDefinitions = relay.NewNodeDefinitions(relay.NodeDefinitionsConfig{
		IDFetcher: func(id string, info graphql.ResolveInfo, ctx context.Context) (interface{}, error) {
			// resolve id from global id
			resolvedID := relay.FromGlobalID(id)

			// based on id and its type, return the object
			switch resolvedID.Type {
			case "Country":
				return getCountryByGlobalId(resolvedID.ID, countryService)
			case "University":
				return getUniversityByGlobalId(resolvedID.ID, universityService)
			default:
				return nil, errors.New("Unknown node type")
			}
		},

		TypeResolve: func(p graphql.ResolveTypeParams) *graphql.Object {
			// based on the type of the value, return GraphQLObjectType
			switch p.Value.(type) {
			case *university.University:
				return universityType
			default:
				return countryType
			}
		},
	})

	countryPropertiesType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "CountryProperties",
		Description: "Properties of a Country object",
		Fields: graphql.Fields{
			"code": &graphql.Field{
				Type:        graphql.String,
				Description: "The code of the country",
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of the country.",
			},
		},
	})

	countryType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Country",
		Description: "A country",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("Country", nil),
			"properties": &graphql.Field{
				Type:        countryPropertiesType,
				Description: "The properties of the country",
			},
		},
		Interfaces: []*graphql.Interface{
			nodeDefinitions.NodeInterface,
		},
	})

	countryMutation := relay.MutationWithClientMutationID(relay.MutationConfig{
		Name: "CreateCountry",
		InputFields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"code": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		OutputFields: graphql.Fields{
			"country": &graphql.Field{
				Type: countryType,
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
			if err := countryService.SaveCountry(&aCountry); err != nil {
				return nil, err
			}
			return map[string]interface{}{
				"country": aCountry,
			}, nil
		},
	})

	addressType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Address",
		Description: "An address",
		Fields: graphql.Fields{
			"line": &graphql.Field{
				Type: graphql.String,
			},
			"city": &graphql.Field{
				Type: graphql.String,
			},
			"state": &graphql.Field{
				Type: graphql.String,
			},
			"postalCode": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	tuitionType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Tuition",
		Description: "A tuition object with amount and source link",
		Fields: graphql.Fields{
			"amount": &graphql.Field{
				Type: graphql.Int,
			},
			"link": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	universityPropertiesType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "UniversityProperties",
		Description: "Properties of an University object",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type:        graphql.String,
				Description: "The name of University.",
			},
			"languages": &graphql.Field{
				Type:        graphql.NewList(graphql.String),
				Description: "The code of the languages used in this University",
			},
			"website": &graphql.Field{
				Type:        graphql.String,
				Description: "The link to the website of this University",
			},
			"programListLink": &graphql.Field{
				Type:        graphql.String,
				Description: "The link to the program list of this University",
			},
			"address": &graphql.Field{
				Type: addressType,
			},
			"tuition": &graphql.Field{
				Type: tuitionType,
			},
		},
	})

	universityType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "University",
		Description: "A University",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("University", nil),
			"properties": &graphql.Field{
				Type:        universityPropertiesType,
				Description: "The properties of the university",
			},
		},
		Interfaces: []*graphql.Interface{
			nodeDefinitions.NodeInterface,
		},
	})

	createUniversityMutation := relay.MutationWithClientMutationID(relay.MutationConfig{
		Name: "CreateUniversity",
		InputFields: graphql.InputObjectConfigFieldMap{
			"countryCode": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"languages": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.NewList(graphql.String)),
			},
			"website": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"programListLink": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"address": &graphql.InputObjectFieldConfig{
				Type: graphql.NewInputObject(graphql.InputObjectConfig{
					Name: "AddressInput",
					Fields: graphql.InputObjectConfigFieldMap{
						"line": &graphql.InputObjectFieldConfig{
							Type: graphql.String,
						},
						"city": &graphql.InputObjectFieldConfig{
							Type: graphql.String,
						},
						"state": &graphql.InputObjectFieldConfig{
							Type: graphql.String,
						},
						"postalCode": &graphql.InputObjectFieldConfig{
							Type: graphql.String,
						},
					},
				}),
			},
			"tuition": &graphql.InputObjectFieldConfig{
				Type: graphql.NewInputObject(graphql.InputObjectConfig{
					Name: "TuitionInput",
					Fields: graphql.InputObjectConfigFieldMap{
						"amount": &graphql.InputObjectFieldConfig{
							Type: graphql.NewNonNull(graphql.Int),
						},
						"link": &graphql.InputObjectFieldConfig{
							Type: graphql.NewNonNull(graphql.String),
						},
					},
				}),
			},
		},

		OutputFields: graphql.Fields{
			"university": &graphql.Field{
				Type: universityType,
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
				address = university.Address{
					Line:       addressMap["line"].(string),
					City:       addressMap["city"].(string),
					State:      addressMap["state"].(string),
					PostalCode: addressMap["postalCode"].(string),
				}
			}

			aUniversity := university.University{
				Name:            name,
				Languages:       languages,
				Website:         website,
				ProgramListLink: programListLink,
				Address:         address,
				Tuition:         tuition,
			}

			if err := universityService.SaveUniversity(&aUniversity, countryCode); err != nil {
				return nil, err
			}

			return map[string]interface{}{
				"university":  aUniversity,
				"countryCode": countryCode,
			}, nil
		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "SabQuery",
		Fields: graphql.Fields{
			"countries": &graphql.Field{
				Name:        "Countries",
				Description: "Get the list of countries supported by the application",
				Type:        graphql.NewList(countryType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return getAllCountries(countryService)
				},
			},

			"country": &graphql.Field{
				Name:        "Country",
				Description: "Get a country by country code",
				Type:        countryType,
				Args: graphql.FieldConfigArgument{
					"code": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.String),
						Description: "The code of the country to get",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return getCountryNodeByCode(p.Args["code"].(string), countryService)
				},
			},

			"universities": &graphql.Field{
				Name:        "Universities",
				Description: "Get the list of universities of this the given country",
				Type:        graphql.NewList(universityType),
				Args: graphql.FieldConfigArgument{
					"countryCode": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.String),
						Description: "The country code of the country from witch to retrieve universities",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return getAllUniversities(p.Args["countryCode"].(string), universityService)
				},
			},

			"university": &graphql.Field{
				Name:        "University",
				Type:        universityType,
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
					return getUniversityNodeByCountryCodeAndUniversityId(p.Args["universityId"].(string), p.Args["countryCode"].(string), universityService)
				},
			},

			"node": nodeDefinitions.NodeField,
		},
	})

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "SabMutation",
		Description: "All the mutation available on Study abroad apy",
		Fields: graphql.Fields{
			"createCountry":    countryMutation,
			"createUniversity": createUniversityMutation,
		},
	})

	aSchema, aErr := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})

	return &aSchema, aErr
}

//func init() {
//
//	/**
//	 * We define a connection between a faction and its ships.
//	 *
//	 * connectionType implements the following type system shorthand:
//	 *   type ShipConnection {
//	 *     edges: [ShipEdge]
//	 *     pageInfo: PageInfo!
//	 *   }
//	 *
//	 * connectionType has an edges field - a list of edgeTypes that implement the
//	 * following type system shorthand:
//	 *   type ShipEdge {
//	 *     cursor: String!
//	 *     node: Ship
//	 *   }
//	 */
//	//shipConnectionDefinition := relay.ConnectionDefinitions(relay.ConnectionConfig{
//	//	Name:     "Ship",
//	//	NodeType: shipType,
//	//})
//	//
//	///**
//	// * We define our faction type, which implements the node interface.
//	// *
//	// * This implements the following type system shorthand:
//	// *   type Faction : Node {
//	// *     id: String!
//	// *     name: String
//	// *     ships: ShipConnection
//	// *   }
//	// */
//	//factionType = graphql.NewObject(graphql.ObjectConfig{
//	//	Name:        "Faction",
//	//	Description: "A faction in the Star Wars saga",
//	//	Fields: graphql.Fields{
//	//		"id": relay.GlobalIDField("Faction", nil),
//	//		"name": &graphql.Field{
//	//			Type:        graphql.String,
//	//			Description: "The name of the faction.",
//	//		},
//	//		"ships": &graphql.Field{
//	//			Type: shipConnectionDefinition.ConnectionType,
//	//			Args: relay.ConnectionArgs,
//	//			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
//	//				// convert args map[string]interface into ConnectionArguments
//	//				args := relay.NewConnectionArguments(p.Args)
//	//
//	//				// get ship objects from current faction
//	//				ships := []interface{}{}
//	//				if faction, ok := p.Source.(*Faction); ok {
//	//					for _, shipId := range faction.Ships {
//	//						ships = append(ships, GetShip(shipId))
//	//					}
//	//				}
//	//				// let relay library figure out the result, given
//	//				// - the list of ships for this faction
//	//				// - and the filter arguments (i.e. first, last, after, before)
//	//				return relay.ConnectionFromArray(ships, args), nil
//	//			},
//	//		},
//	//	},
//	//	Interfaces: []*graphql.Interface{
//	//		nodeDefinitions.NodeInterface,
//	//	},
//	//})
//
//	/**
//	 * This is the type that will be the root of our query, and the
//	 * entry point into our schema.
//	 *
//	 * This implements the following type system shorthand:
//	 *   type Query {
//	 *     rebels: Faction
//	 *     empire: Faction
//	 *     node(id: String!): Node
//	 *   }
//	 */
//	//queryType := graphql.NewObject(graphql.ObjectConfig{
//	//	Name: "Query",
//	//	Fields: graphql.Fields{
//	//		"rebels": &graphql.Field{
//	//			Type: factionType,
//	//			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
//	//				return GetRebels(), nil
//	//			},
//	//		},
//	//		"empire": &graphql.Field{
//	//			Type: factionType,
//	//			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
//	//				return GetEmpire(), nil
//	//			},
//	//		},
//	//		"node": nodeDefinitions.NodeField,
//	//	},
//	//})
//
//	/**
//	 * This will return a GraphQLField for our ship
//	 * mutation.
//	 *
//	 * It creates these two types implicitly:
//	 *   input IntroduceShipInput {
//	 *     clientMutationID: string!
//	 *     shipName: string!
//	 *     factionId: ID!
//	 *   }
//	 *
//	 *   input IntroduceShipPayload {
//	 *     clientMutationID: string!
//	 *     ship: Ship
//	 *     faction: Faction
//	 *   }
//	 */
//	shipMutation := relay.MutationWithClientMutationID(relay.MutationConfig{
//		Name: "IntroduceShip",
//		InputFields: graphql.InputObjectConfigFieldMap{
//			"shipName": &graphql.InputObjectFieldConfig{
//				Type: graphql.NewNonNull(graphql.String),
//			},
//			"factionId": &graphql.InputObjectFieldConfig{
//				Type: graphql.NewNonNull(graphql.ID),
//			},
//		},
//		OutputFields: graphql.Fields{
//			"ship": &graphql.Field{
//				Type: shipType,
//				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
//					if payload, ok := p.Source.(map[string]interface{}); ok {
//						return GetShip(payload["shipId"].(string)), nil
//					}
//					return nil, nil
//				},
//			},
//			"faction": &graphql.Field{
//				Type: factionType,
//				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
//					if payload, ok := p.Source.(map[string]interface{}); ok {
//						return GetFaction(payload["factionId"].(string)), nil
//					}
//					return nil, nil
//				},
//			},
//		},
//		MutateAndGetPayload: func(inputMap map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
//			// `inputMap` is a map with keys/fields as specified in `InputFields`
//			// Note, that these fields were specified as non-nullables, so we can assume that it exists.
//			shipName := inputMap["shipName"].(string)
//			factionId := inputMap["factionId"].(string)
//
//			// This mutation involves us creating (introducing) a new ship
//			newShip := CreateShip(shipName, factionId)
//			// return payload
//			return map[string]interface{}{
//				"shipId":    newShip.ID,
//				"factionId": factionId,
//			}, nil
//		},
//	})
//
//	/**
//	 * This is the type that will be the root of our mutations, and the
//	 * entry point into performing writes in our schema.
//	 *
//	 * This implements the following type system shorthand:
//	 *   type Mutation {
//	 *     introduceShip(input IntroduceShipInput!): IntroduceShipPayload
//	 *   }
//	 */
//
//	mutationType := graphql.NewObject(graphql.ObjectConfig{
//		Name: "Mutation",
//		Fields: graphql.Fields{
//			"introduceShip": shipMutation,
//		},
//	})
//
//	/**
//	 * Finally, we construct our schema (whose starting query type is the query
//	 * type we defined above) and export it.
//	 */
//	var err error
//	schema, err = graphql.NewSchema(graphql.SchemaConfig{
//		Query:    queryType,
//		Mutation: mutationType,
//	})
//	if err != nil {
//		// panic if there is an error in schema
//		panic(err)
//	}
//}
