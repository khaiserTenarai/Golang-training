package graphql

import gql "github.com/graphql-go/graphql"

func newEmployeeType() *gql.Object {
	return gql.NewObject(gql.ObjectConfig{
		Name: "Employee",
		Fields: gql.Fields{
			"id":         &gql.Field{Type: gql.NewNonNull(gql.ID)},
			"name":       &gql.Field{Type: gql.NewNonNull(gql.String)},
			"email":      &gql.Field{Type: gql.NewNonNull(gql.String)},
			"department": &gql.Field{Type: gql.NewNonNull(gql.String)},
			"salary":     &gql.Field{Type: gql.NewNonNull(gql.Float)},
		},
	})
}
