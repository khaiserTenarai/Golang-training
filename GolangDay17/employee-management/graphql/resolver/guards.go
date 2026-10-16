package resolver

import (
	"employee-management/auth"
	gql "github.com/graphql-go/graphql"
)


func LoginRequired(next gql.FieldResolveFn) gql.FieldResolveFn {
	return func(p gql.ResolveParams) (interface{}, error) {
		if _, err := auth.RequireUser(p.Context); err != nil {
			return nil, toGraphQLError(err, "")
		}
		return next(p)
	}
}


func AdminOnly(next gql.FieldResolveFn) gql.FieldResolveFn {
	return func(p gql.ResolveParams) (interface{}, error) {
		if _, err := auth.RequireRole(p.Context, auth.RoleAdmin); err != nil {
			return nil, toGraphQLError(err, "")
		}
		return next(p)
	}
}
