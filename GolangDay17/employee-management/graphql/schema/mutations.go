package schema

import (
	"employee-management/graphql/resolver"
	gql "github.com/graphql-go/graphql"
)

// mutationType lists everything a client can CHANGE.
// login is public, everything else needs the ADMIN role.
func mutationType(r *resolver.Resolver, t *objectTypes, in *inputTypes) *gql.Object {
	idArg := gql.FieldConfigArgument{"id": &gql.ArgumentConfig{Type: gql.NewNonNull(gql.ID)}}

	return gql.NewObject(gql.ObjectConfig{Name: "Mutation", Fields: gql.Fields{
		"login": &gql.Field{
			Description: "Log in and receive a token",
			Type:        gql.NewNonNull(t.AuthPayload),
			Args: gql.FieldConfigArgument{
				"username": &gql.ArgumentConfig{Type: gql.NewNonNull(gql.String)},
				"password": &gql.ArgumentConfig{Type: gql.NewNonNull(gql.String)},
			},
			Resolve: r.Login, // public
		},
		"createDepartment": &gql.Field{
			Type:    gql.NewNonNull(t.Department),
			Args:    gql.FieldConfigArgument{"input": &gql.ArgumentConfig{Type: gql.NewNonNull(in.DepartmentInput)}},
			Resolve: resolver.AdminOnly(r.CreateDepartment),
		},
		"createEmployee": &gql.Field{
			Type:    gql.NewNonNull(t.Employee),
			Args:    gql.FieldConfigArgument{"input": &gql.ArgumentConfig{Type: gql.NewNonNull(in.EmployeeInput)}},
			Resolve: resolver.AdminOnly(r.CreateEmployee),
		},
		"updateEmployee": &gql.Field{
			Type: gql.NewNonNull(t.Employee),
			Args: gql.FieldConfigArgument{
				"id":    &gql.ArgumentConfig{Type: gql.NewNonNull(gql.ID)},
				"input": &gql.ArgumentConfig{Type: gql.NewNonNull(in.EmployeeInput)},
			},
			Resolve: resolver.AdminOnly(r.UpdateEmployee),
		},
		"deleteEmployee": &gql.Field{
			Type:    gql.NewNonNull(gql.Boolean),
			Args:    idArg,
			Resolve: resolver.AdminOnly(r.DeleteEmployee),
		},
	}})
}
