package graphql

import gql "github.com/graphql-go/graphql"

var EmployeeType = gql.NewObject(gql.ObjectConfig{
	Name: "Employee",
	Fields: gql.Fields{
		"id":         &gql.Field{Type: gql.Int},
		"name":       &gql.Field{Type: gql.String},
		"email":      &gql.Field{Type: gql.String},
		"department": &gql.Field{Type: gql.String},
		"salary":     &gql.Field{Type: gql.Float},
	},
})

var EmployeeInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "EmployeeInput",
	Fields: gql.InputObjectConfigFieldMap{
		"name":       &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
		"email":      &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
		"department": &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
		"salary":     &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.Float)},
	},
})
