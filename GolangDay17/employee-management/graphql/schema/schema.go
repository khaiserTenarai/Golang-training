// Package schema builds the GraphQL schema from small modules:
//   types.go     - what the API returns
//   inputs.go    - what clients send (input types and enums)
//   queries.go   - read operations
//   mutations.go - write operations
package schema

import (
	"employee-management/graphql/resolver"
	gql "github.com/graphql-go/graphql"
)

// Build puts all the modules together.
func Build(r *resolver.Resolver) (gql.Schema, error) {
	types := newObjectTypes(r)
	inputs := newInputTypes()
	return gql.NewSchema(gql.SchemaConfig{
		Query:    queryType(r, types, inputs),
		Mutation: mutationType(r, types, inputs),
	})
}
