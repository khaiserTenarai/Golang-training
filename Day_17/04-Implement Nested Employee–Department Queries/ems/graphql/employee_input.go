package graphql

import gql "github.com/graphql-go/graphql"

func newEmployeeInputType() *gql.InputObject {
	return gql.NewInputObject(gql.InputObjectConfig{
		Name: "EmployeeInput",
		Fields: gql.InputObjectConfigFieldMap{
			"name":       &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
			"email":      &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
			"department": &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
			"salary":     &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.Float)},
		},
	})
}
