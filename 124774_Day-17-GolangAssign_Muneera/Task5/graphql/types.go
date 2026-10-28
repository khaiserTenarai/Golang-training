package graphql

import gql "github.com/graphql-go/graphql"

var DepartmentType = gql.NewObject(gql.ObjectConfig{
	Name: "Department",
	Fields: gql.Fields{
		"id":   &gql.Field{Type: gql.Int},
		"name": &gql.Field{Type: gql.String},
	},
})

var EmployeeType = gql.NewObject(gql.ObjectConfig{
	Name: "Employee",
	Fields: gql.Fields{
		"id":           &gql.Field{Type: gql.Int},
		"name":         &gql.Field{Type: gql.String},
		"email":        &gql.Field{Type: gql.String},
		"departmentId": &gql.Field{Type: gql.Int},
		"department":   &gql.Field{Type: DepartmentType},
		"salary":       &gql.Field{Type: gql.Float},
	},
})

var EmployeeInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "EmployeeInput",
	Fields: gql.InputObjectConfigFieldMap{
		"name": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"email": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.String),
		},
		"departmentId": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.Int),
		},
		"salary": &gql.InputObjectFieldConfig{
			Type: gql.NewNonNull(gql.Float),
		},
	},
})
