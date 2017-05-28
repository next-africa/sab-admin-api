package graphql

import (
	"github.com/graphql-go/graphql"
	"sab.com/domain/university"
)

type UniversityNode struct {
	Id         string                 `json:"id"`
	Properties *university.University `json:"properties"`
}

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

var AddressInput = graphql.NewInputObject(graphql.InputObjectConfig{
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
})

var TuitionInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "TuitionInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"amount": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"link": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})

var fullUniversityInputObjectConfigMap = graphql.InputObjectConfigFieldMap{
	"id": &graphql.InputObjectFieldConfig{
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
		Type: AddressInput,
	},
	"tuition": &graphql.InputObjectFieldConfig{
		Type: TuitionInput,
	},
}

func getCreateUniversityInputType() *graphql.InputObject {
	createUniversityInputObjectConfigMap := graphql.InputObjectConfigFieldMap{}
	for k, v := range fullUniversityInputObjectConfigMap {
		createUniversityInputObjectConfigMap[k] = v
	}
	delete(createUniversityInputObjectConfigMap, "id")

	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name:   "UniversityInputForCreate",
		Fields: createUniversityInputObjectConfigMap,
	})
}

func getUpdateUniversityInputType() *graphql.InputObject {
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name:   "UniversityInputForUpdate",
		Fields: fullUniversityInputObjectConfigMap,
	})
}

var CreateUniversityInputFields = graphql.InputObjectConfigFieldMap{
	"countryCode": &graphql.InputObjectFieldConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"university": &graphql.InputObjectFieldConfig{
		Type: graphql.NewNonNull(getCreateUniversityInputType()),
	},
}

var UpdateUniversityInputFields = graphql.InputObjectConfigFieldMap{
	"countryCode": &graphql.InputObjectFieldConfig{
		Type: graphql.NewNonNull(graphql.String),
	},
	"university": &graphql.InputObjectFieldConfig{
		Type: graphql.NewNonNull(getUpdateUniversityInputType()),
	},
}
