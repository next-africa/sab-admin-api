package graphql

import (
	"github.com/graphql-go/graphql"
	"sab.com/domain/country"
)

type CountryNode struct {
	Id         string           `json:"id"`
	Properties *country.Country `json:"properties"`
}

var CountryPropertiesType = graphql.NewObject(graphql.ObjectConfig{
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

var CreateCountryInputFields = graphql.InputObjectConfigFieldMap{
	"name": &graphql.InputObjectFieldConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"code": &graphql.InputObjectFieldConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
}