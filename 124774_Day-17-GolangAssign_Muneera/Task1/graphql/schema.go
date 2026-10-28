package graphql

import (
	gql "github.com/graphql-go/graphql"

	"employee-management/service"
)

type Server struct {
	schema gql.Schema
}

func NewServer(svc service.EmployeeService) (*Server, error) {
	// Build the query module.
	queryType := NewQueryType(svc)

	// Build the mutation module.
	mutationType := NewMutationType(svc)

	// Combine both modules into one GraphQL schema.
	schema, err := gql.NewSchema(gql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
	if err != nil {
		return nil, err
	}

	return &Server{schema: schema}, nil
}

func (s *Server) Schema() gql.Schema {
	return s.schema
}
