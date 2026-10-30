package graphql

import (
	"fmt"

	"day18/service"

	gql "github.com/graphql-go/graphql"
)

type Server struct {
	service service.EmployeeService
	schema  gql.Schema
}

func NewServer(s service.EmployeeService) (*Server, error) {
	srv := &Server{service: s}
	sc, err := srv.buildSchema()
	if err != nil {
		return nil, err
	}
	srv.schema = sc
	return srv, nil
}

func (s *Server) buildSchema() (gql.Schema, error) {
	empType := gql.NewObject(gql.ObjectConfig{
		Name: "Employee",
		Fields: gql.Fields{
			"id":            &gql.Field{Type: gql.NewNonNull(gql.Int)},
			"first_name":    &gql.Field{Type: gql.NewNonNull(gql.String)},
			"last_name":     &gql.Field{Type: gql.NewNonNull(gql.String)},
			"email":         &gql.Field{Type: gql.NewNonNull(gql.String)},
			"department_id": &gql.Field{Type: gql.NewNonNull(gql.Int)},
			"status":        &gql.Field{Type: gql.NewNonNull(gql.String)},
		},
	})

	inputParam := gql.NewInputObject(gql.InputObjectConfig{
		Name: "EmployeeInput",
		Fields: gql.InputObjectConfigFieldMap{
			"first_name":    &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
			"last_name":     &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
			"email":         &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
			"department_id": &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.Int)},
		},
	})

	query := gql.NewObject(gql.ObjectConfig{
		Name: "Query",
		Fields: gql.Fields{
			"employee": &gql.Field{
				Type: empType,
				Args: gql.FieldConfigArgument{
					"id": &gql.ArgumentConfig{Type: gql.NewNonNull(gql.Int)},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					id := int64(p.Args["id"].(int))
					return s.service.GetEmployee(p.Context, id)
				},
			},
		},
	})

	mutation := gql.NewObject(gql.ObjectConfig{
		Name: "Mutation",
		Fields: gql.Fields{
			"createEmployee": &gql.Field{
				Type: empType,
				Args: gql.FieldConfigArgument{
					"input": &gql.ArgumentConfig{Type: gql.NewNonNull(inputParam)},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					input := p.Args["input"].(map[string]interface{})

					firstName := fmt.Sprint(input["first_name"])
					lastName := fmt.Sprint(input["last_name"])
					email := fmt.Sprint(input["email"])
					deptID := int64(input["department_id"].(int))

					return s.service.OnboardEmployee(p.Context, firstName, lastName, email, deptID)
				},
			},
		},
	})

	return gql.NewSchema(gql.SchemaConfig{
		Query:    query,
		Mutation: mutation,
	})
}
