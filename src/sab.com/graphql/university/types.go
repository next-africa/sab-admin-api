package graphql

import "github.com/next-africa/graphql-go"

var AddressType = graphql.NewObject(graphql.ObjectConfig{
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

var TuitionType = graphql.NewObject(graphql.ObjectConfig{
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

var UniversityPropertiesType = graphql.NewObject(graphql.ObjectConfig{
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
			Type: AddressType,
		},
		"tuition": &graphql.Field{
			Type: TuitionType,
		},
	},
})

var CreateUniversityInputFields = graphql.InputObjectConfigFieldMap{
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
}
