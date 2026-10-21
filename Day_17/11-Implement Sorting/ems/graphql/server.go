package graphql

import (
	"fmt"

	"employee-management/service"

	gql "github.com/graphql-go/graphql"
)

type Server struct {
	service service.EmployeeService
	schema  gql.Schema
}

func NewServer(s service.EmployeeService) (*Server, error) {
	server := &Server{service: s}

	employeeType := newEmployeeType()
	inputType := newEmployeeInputType()

	query := newQueryType(server, employeeType)
	mutation := newMutationType(server, employeeType, inputType)

	schema, err := gql.NewSchema(gql.SchemaConfig{
		Query: query, Mutation: mutation,
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
