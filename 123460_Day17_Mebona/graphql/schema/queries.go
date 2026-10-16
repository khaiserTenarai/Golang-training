package schema

import (
	"employee-management/graphql/resolver"
	gql "github.com/graphql-go/graphql"
)

// queryType lists everything a client can READ.
// LoginRequired(...) means: "you must be logged in to use this field".
func queryType(r *resolver.Resolver, t *objectTypes, in *inputTypes) *gql.Object {
	idArg := gql.FieldConfigArgument{"id": &gql.ArgumentConfig{Type: gql.NewNonNull(gql.ID)}}

	return gql.NewObject(gql.ObjectConfig{Name: "Query", Fields: gql.Fields{
		"me": &gql.Field{
			Description: "The logged-in user",
			Type:        gql.NewNonNull(t.CurrentUser),
			Resolve:     resolver.LoginRequired(r.Me),
		},
		"employee": &gql.Field{
			Description: "Get one employee by id",
			Type:        t.Employee,
			Args:        idArg,
			Resolve:     resolver.LoginRequired(r.Employee),
		},
		"employees": &gql.Field{
			Description: "List employees with pagination, filtering and sorting",
			Type:        gql.NewNonNull(t.EmployeePage),
			Args: gql.FieldConfigArgument{
				// pagination
				"page":     &gql.ArgumentConfig{Type: gql.Int, Description: "Starts at 1 (default 1)"},
				"pageSize": &gql.ArgumentConfig{Type: gql.Int, Description: "1 to 100 (default 10)"},
				// filtering
				"search":       &gql.ArgumentConfig{Type: gql.String, Description: "Matches name or email"},
				"departmentId": &gql.ArgumentConfig{Type: gql.ID, Description: "Only this department"},
				"minSalary":    &gql.ArgumentConfig{Type: gql.Float},
				"maxSalary":    &gql.ArgumentConfig{Type: gql.Float},
				// sorting
				"sortBy":    &gql.ArgumentConfig{Type: in.SortField, Description: "Default ID"},
				"sortOrder": &gql.ArgumentConfig{Type: in.SortOrder, Description: "Default ASC"},
			},
			Resolve: resolver.LoginRequired(r.Employees),
		},
		"department": &gql.Field{
			Description: "Get one department by id",
			Type:        t.Department,
			Args:        idArg,
			Resolve:     resolver.LoginRequired(r.Department),
		},
		"departments": &gql.Field{
			Type:    gql.NewNonNull(gql.NewList(gql.NewNonNull(t.Department))),
			Resolve: resolver.LoginRequired(r.Departments),
		},
	}})
}
