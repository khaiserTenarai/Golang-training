package graphql

import (
	"fmt"

	"employee-management/graphql/resolver"
	"employee-management/service"

	gql "github.com/graphql-go/graphql"
)

type Server struct {
	resolvers *resolver.EmployeeResolver
	schema    gql.Schema
}

func NewServer(s service.EmployeeService) (*Server, error) {
	server := &Server{
		resolvers: resolver.NewEmployeeResolver(s),
	}

	employeeType := newEmployeeType()
	employeeInput := newEmployeeInputType()

	queryType := newQueryType(server, employeeType)
	mutationType := newMutationType(
		server,
		employeeType,
		employeeInput,
	)

	schema, err := gql.NewSchema(gql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
	if err != nil {
		return nil, fmt.Errorf("create GraphQL schema: %w", err)
	}

	server.schema = schema
	return server, nil
}

func (s *Server) Schema() *gql.Schema {
	return &s.schema
}
