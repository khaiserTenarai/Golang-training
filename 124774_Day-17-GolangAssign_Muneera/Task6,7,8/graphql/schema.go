package graphql

import (
	"fmt"
	"strings"

	"employee-management/middleware"
	"employee-management/model"
	"employee-management/service"

	gql "github.com/graphql-go/graphql"
)

type Server struct {
	service service.EmployeeService
	schema  gql.Schema
}

func NewServer(s service.EmployeeService) (*Server, error) {
	server := &Server{service: s}

	schema, err := server.buildSchema()
	if err != nil {
		return nil, err
	}

	server.schema = schema
	return server, nil
}

func (x *Server) buildSchema() (gql.Schema, error) {
	employeeType := gql.NewObject(gql.ObjectConfig{
		Name: "Employee",
		Fields: gql.Fields{
			"id":         &gql.Field{Type: gql.NewNonNull(gql.Int)},
			"name":       &gql.Field{Type: gql.NewNonNull(gql.String)},
			"email":      &gql.Field{Type: gql.NewNonNull(gql.String)},
			"department": &gql.Field{Type: gql.NewNonNull(gql.String)},
			"salary":     &gql.Field{Type: gql.NewNonNull(gql.Float)},
		},
	})

	employeeInput := gql.NewInputObject(gql.InputObjectConfig{
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
				Type: employeeType,
				Args: gql.FieldConfigArgument{
					"id": &gql.ArgumentConfig{
						Type: gql.NewNonNull(gql.Int),
					},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if !ok || id <= 0 {
						return nil, fmt.Errorf("employee id must be a positive integer")
					}

					employee, err := x.service.GetByID(uint(id))
					if service.IsNotFound(err) {
						return nil, fmt.Errorf("employee not found")
					}

					return employee, err
				},
			},

			"employees": &gql.Field{
				Type: gql.NewList(employeeType),
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					return x.service.GetAll()
				},
			},
		},
	})

	mutation := gql.NewObject(gql.ObjectConfig{
		Name: "Mutation",
		Fields: gql.Fields{
			"createEmployee": &gql.Field{
				Type: employeeType,
				Args: gql.FieldConfigArgument{
					"input": &gql.ArgumentConfig{
						Type: gql.NewNonNull(employeeInput),
					},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					if err := middleware.RequireAdmin(p.Context); err != nil {
						return nil, err
					}

					input, ok := p.Args["input"].(map[string]interface{})
					if !ok {
						return nil, fmt.Errorf("invalid employee input")
					}

					employee := fromInput(input)

					if err := x.service.Create(&employee); err != nil {
						return nil, err
					}

					return employee, nil
				},
			},

			"updateEmployee": &gql.Field{
				Type: employeeType,
				Args: gql.FieldConfigArgument{
					"id": &gql.ArgumentConfig{
						Type: gql.NewNonNull(gql.Int),
					},
					"input": &gql.ArgumentConfig{
						Type: gql.NewNonNull(employeeInput),
					},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					if err := middleware.RequireAdmin(p.Context); err != nil {
						return nil, err
					}

					id, ok := p.Args["id"].(int)
					if !ok || id <= 0 {
						return nil, fmt.Errorf("employee id must be a positive integer")
					}

					input, ok := p.Args["input"].(map[string]interface{})
					if !ok {
						return nil, fmt.Errorf("invalid employee input")
					}

					employee := fromInput(input)
					return x.service.Update(uint(id), &employee)
				},
			},

			"deleteEmployee": &gql.Field{
				Type: gql.NewNonNull(gql.Boolean),
				Args: gql.FieldConfigArgument{
					"id": &gql.ArgumentConfig{
						Type: gql.NewNonNull(gql.Int),
					},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					if err := middleware.RequireAdmin(p.Context); err != nil {
						return false, err
					}

					id, ok := p.Args["id"].(int)
					if !ok || id <= 0 {
						return false, fmt.Errorf("employee id must be a positive integer")
					}

					if err := x.service.Delete(uint(id)); err != nil {
						if service.IsNotFound(err) {
							return false, fmt.Errorf("employee not found")
						}

						return false, err
					}

					return true, nil
				},
			},
		},
	})

	return gql.NewSchema(gql.SchemaConfig{
		Query:    query,
		Mutation: mutation,
	})
}

func fromInput(input map[string]interface{}) model.Employee {
	name, _ := input["name"].(string)
	email, _ := input["email"].(string)
	department, _ := input["department"].(string)

	salary, _ := input["salary"].(float64)

	return model.Employee{
		Name:       strings.TrimSpace(name),
		Email:      strings.TrimSpace(email),
		Department: strings.TrimSpace(department),
		Salary:     salary,
	}
}
