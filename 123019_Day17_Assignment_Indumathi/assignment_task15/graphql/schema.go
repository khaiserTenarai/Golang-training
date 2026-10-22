package graphql

import (
	"errors"
	"fmt"

	"employee-management/model"
	"employee-management/service"

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
			"id":         &gql.Field{Type: gql.NewNonNull(gql.Int)},
			"name":       &gql.Field{Type: gql.NewNonNull(gql.String)},
			"email":      &gql.Field{Type: gql.NewNonNull(gql.String)},
			"department": &gql.Field{Type: gql.NewNonNull(gql.String)},
			"salary":     &gql.Field{Type: gql.NewNonNull(gql.Float)},
		},
	})

	inputParam := gql.NewInputObject(gql.InputObjectConfig{
		Name: "EmployeeInput",
		Fields: gql.InputObjectConfigFieldMap{
			"name":       &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
			"email":      &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
			"department": &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
			"salary":     &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.Float)},
		},
	})

	query := gql.NewObject(gql.ObjectConfig{
		Name: "Query",
		Fields: gql.Fields{
			"employee": &gql.Field{
				Type: empType,
				Args: gql.FieldConfigArgument{"id": &gql.ArgumentConfig{Type: gql.NewNonNull(gql.Int)}},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					id := int64(p.Args["id"].(int))
					return s.service.GetByID(p.Context, id)
				},
			},
			"employees": &gql.Field{
				Type: gql.NewList(empType),
				Args: gql.FieldConfigArgument{"sort": &gql.ArgumentConfig{Type: gql.String}},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					sortBy, _ := p.Args["sort"].(string)
					return s.service.GetAll(p.Context, sortBy)
				},
			},
		},
	})

	mutation := gql.NewObject(gql.ObjectConfig{
		Name: "Mutation",
		Fields: gql.Fields{
			"createEmployee": &gql.Field{
				Type: empType,
				Args: gql.FieldConfigArgument{"input": &gql.ArgumentConfig{Type: gql.NewNonNull(inputParam)}},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					emp := fromInput(p.Args["input"].(map[string]interface{}))
					if err := s.service.Create(p.Context, &emp); err != nil {
						return nil, err
					}
					return emp, nil
				},
			},
			"updateEmployee": &gql.Field{
				Type: empType,
				Args: gql.FieldConfigArgument{
					"id":    &gql.ArgumentConfig{Type: gql.NewNonNull(gql.Int)},
					"input": &gql.ArgumentConfig{Type: gql.NewNonNull(inputParam)},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					id := int64(p.Args["id"].(int))
					emp := fromInput(p.Args["input"].(map[string]interface{}))
					return s.service.Update(p.Context, id, &emp)
				},
			},
			"deleteEmployee": &gql.Field{
				Type: gql.NewNonNull(gql.Boolean),
				Args: gql.FieldConfigArgument{"id": &gql.ArgumentConfig{Type: gql.NewNonNull(gql.Int)}},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					id := int64(p.Args["id"].(int))
					if err := s.service.Delete(p.Context, id); err != nil {
						return false, errors.New("failed to delete employee")
					}
					return true, nil
				},
			},
		},
	})

	return gql.NewSchema(gql.SchemaConfig{Query: query, Mutation: mutation})
}

func fromInput(m map[string]interface{}) model.Employee {
	return model.Employee{
		Name:       fmt.Sprint(m["name"]),
		Email:      fmt.Sprint(m["email"]),
		Department: fmt.Sprint(m["department"]),
		Salary:     m["salary"].(float64),
	}
}