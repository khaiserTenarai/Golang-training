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
	server := &Server{}

	schema, err := server.buildSchema(svc)
	if err != nil {
		return nil, err
	}

	server.schema = schema
	return server, nil
}

func (s *Server) buildSchema(
	svc service.EmployeeService,
) (gql.Schema, error) {
	query := gql.NewObject(gql.ObjectConfig{
		Name: "Query",
		Fields: gql.Fields{
			"employee": &gql.Field{
				Type: EmployeeType,
				Args: gql.FieldConfigArgument{
					"id": &gql.ArgumentConfig{
						Type: gql.NewNonNull(gql.Int),
					},
				},
				Resolve: resolvers.EmployeeByID(svc),
			},
			"employees": &gql.Field{
				Type:    gql.NewList(EmployeeType),
				Resolve: resolvers.Employees(svc),
			},
		},
	})

	mutation := gql.NewObject(gql.ObjectConfig{
		Name: "Mutation",
		Fields: gql.Fields{
			"createEmployee": &gql.Field{
				Type: EmployeeType,
				Args: gql.FieldConfigArgument{
					"input": &gql.ArgumentConfig{
						Type: gql.NewNonNull(EmployeeInputType),
					},
				},
				Resolve: resolvers.CreateEmployee(svc),
			},
			"updateEmployee": &gql.Field{
				Type: EmployeeType,
				Args: gql.FieldConfigArgument{
					"id": &gql.ArgumentConfig{
						Type: gql.NewNonNull(gql.Int),
					},
					"input": &gql.ArgumentConfig{
						Type: gql.NewNonNull(EmployeeInputType),
					},
				},
				Resolve: resolvers.UpdateEmployee(svc),
			},
			"deleteEmployee": &gql.Field{
				Type: gql.NewNonNull(gql.Boolean),
				Args: gql.FieldConfigArgument{
					"id": &gql.ArgumentConfig{
						Type: gql.NewNonNull(gql.Int),
					},
				},
				Resolve: resolvers.DeleteEmployee(svc),
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
