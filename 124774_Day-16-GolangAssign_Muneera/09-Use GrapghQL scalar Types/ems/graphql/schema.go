package graphql

import (
	"fmt"
	"strconv"

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

	employeeType := gql.NewObject(gql.ObjectConfig{
		Name: "Employee",
		Fields: gql.Fields{
			"id":         &gql.Field{Type: gql.NewNonNull(gql.ID)},
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

	queryType := gql.NewObject(gql.ObjectConfig{
		Name: "Query",
		Fields: gql.Fields{
			"employees": &gql.Field{
				Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(employeeType))),
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					return server.service.GetAll()
				},
			},
			"employee": &gql.Field{
				Type: employeeType,
				Args: gql.FieldConfigArgument{
					"id": &gql.ArgumentConfig{Type: gql.NewNonNull(gql.ID)},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					id, err := parseGraphQLID(p.Args["id"])
					if err != nil {
						return nil, err
					}

					return server.service.GetByID(id)
				},
			},
		},
	})

	mutationType := gql.NewObject(gql.ObjectConfig{
		Name: "Mutation",
		Fields: gql.Fields{
			"createEmployee": &gql.Field{
				Type: employeeType,
				Args: gql.FieldConfigArgument{
					"input": &gql.ArgumentConfig{Type: gql.NewNonNull(employeeInput)},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					employee, err := employeeFromInput(p.Args["input"])
					if err != nil {
						return nil, err
					}

					if err := server.service.Create(&employee); err != nil {
						return nil, err
					}

					return &employee, nil
				},
			},
			"updateEmployee": &gql.Field{
				Type: employeeType,
				Args: gql.FieldConfigArgument{
					"id":    &gql.ArgumentConfig{Type: gql.NewNonNull(gql.ID)},
					"input": &gql.ArgumentConfig{Type: gql.NewNonNull(employeeInput)},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					id, err := parseGraphQLID(p.Args["id"])
					if err != nil {
						return nil, err
					}

					employee, err := employeeFromInput(p.Args["input"])
					if err != nil {
						return nil, err
					}

					return server.service.Update(id, employee)
				},
			},
			"deleteEmployee": &gql.Field{
				Type: gql.NewNonNull(gql.Boolean),
				Args: gql.FieldConfigArgument{
					"id": &gql.ArgumentConfig{Type: gql.NewNonNull(gql.ID)},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					id, err := parseGraphQLID(p.Args["id"])
					if err != nil {
						return nil, err
					}

					if err := server.service.Delete(id); err != nil {
						return nil, err
					}

					return true, nil
				},
			},
		},
	})

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

// Schema returns the schema used by the HTTP handler.
func (s *Server) Schema() *gql.Schema {
	return &s.schema
}

func parseGraphQLID(value interface{}) (uint, error) {
	var id uint64
	var err error

	switch v := value.(type) {
	case string:
		id, err = strconv.ParseUint(v, 10, 64)
	case int:
		if v <= 0 {
			return 0, fmt.Errorf("ID must be a positive integer")
		}
		id = uint64(v)
	case uint:
		id = uint64(v)
	case uint64:
		id = v
	case float64:
		if v <= 0 || v != float64(uint64(v)) {
			return 0, fmt.Errorf("ID must be a positive integer")
		}
		id = uint64(v)
	default:
		return 0, fmt.Errorf("invalid employee ID")
	}

	if err != nil || id == 0 {
		return 0, fmt.Errorf("ID must be a positive integer")
	}

	return uint(id), nil
}

func employeeFromInput(value interface{}) (model.Employee, error) {
	input, ok := value.(map[string]interface{})
	if !ok {
		return model.Employee{}, fmt.Errorf("invalid employee input")
	}

	name, okName := input["name"].(string)
	email, okEmail := input["email"].(string)
	department, okDepartment := input["department"].(string)

	if !okName || !okEmail || !okDepartment {
		return model.Employee{}, fmt.Errorf("name, email, and department are required")
	}

	salary, ok := input["salary"].(float64)
	if !ok {
		return model.Employee{}, fmt.Errorf("salary must be a number")
	}

	return model.Employee{
		Name:       name,
		Email:      email,
		Department: department,
		Salary:     salary,
	}, nil
}