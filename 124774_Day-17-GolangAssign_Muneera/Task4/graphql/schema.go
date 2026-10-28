package graphql

import (
	"fmt"

	"employee-management/graphql/resolvers"
	"employee-management/service"
	gql "github.com/graphql-go/graphql"
)

type Server struct {
	schema gql.Schema
}

func NewServer(svc service.EmployeeService) (*Server, error) {
	// Inject the application service into the resolver once.
	employeeResolver := resolvers.NewEmployeeResolver(svc)

	server := &Server{}
	schema, err := server.buildSchema(employeeResolver)
	if err != nil {
		return nil, err
	}
	server.schema = schema
	return server, nil
}

func (s *Server) buildSchema(r *resolvers.EmployeeResolver) (gql.Schema, error) {
	query := gql.NewObject(gql.ObjectConfig{
		Name: "Query",
		Fields: gql.Fields{
			"employee": &gql.Field{
				Type: EmployeeType,
				Args: gql.FieldConfigArgument{
					"id": &gql.ArgumentConfig{Type: gql.NewNonNull(gql.Int)},
				},
				Resolve: r.EmployeeByID,
			},
			"employees": &gql.Field{
				Type:    gql.NewList(EmployeeType),
				Resolve: r.Employees,
			},
		},
	})

	mutation := gql.NewObject(gql.ObjectConfig{
		Name: "Mutation",
		Fields: gql.Fields{
			"createEmployee": &gql.Field{
				Type: EmployeeType,
				Args: gql.FieldConfigArgument{
					"input": &gql.ArgumentConfig{Type: gql.NewNonNull(EmployeeInputType)},
				},
				Resolve: r.CreateEmployee,
			},
			"updateEmployee": &gql.Field{
				Type: EmployeeType,
				Args: gql.FieldConfigArgument{
					"id":    &gql.ArgumentConfig{Type: gql.NewNonNull(gql.Int)},
					"input": &gql.ArgumentConfig{Type: gql.NewNonNull(EmployeeInputType)},
				},
				Resolve: r.UpdateEmployee,
			},
			"deleteEmployee": &gql.Field{
				Type: gql.NewNonNull(gql.Boolean),
				Args: gql.FieldConfigArgument{
					"id": &gql.ArgumentConfig{Type: gql.NewNonNull(gql.Int)},
				},
				Resolve: r.DeleteEmployee,
			},
		},
	})

	schema, err := gql.NewSchema(gql.SchemaConfig{
		Query:    query,
		Mutation: mutation,
	})
	if err != nil {
		return gql.Schema{}, fmt.Errorf("create GraphQL schema: %w", err)
	}
	return schema, nil
}
