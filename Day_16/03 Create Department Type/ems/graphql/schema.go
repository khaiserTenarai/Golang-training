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

	// Department Type
	departmentType := gql.NewObject(gql.ObjectConfig{
		Name: "Department",
		Fields: gql.Fields{
			"name": &gql.Field{
				Type: gql.String,
			},
		},
	})

	// Employee Type: normal department + Department Type
	employeeType := gql.NewObject(gql.ObjectConfig{
		Name: "Employee",
		Fields: gql.Fields{
			"id":     &gql.Field{Type: gql.ID},
			"name":   &gql.Field{Type: gql.String},
			"email":  &gql.Field{Type: gql.String},
			"salary": &gql.Field{Type: gql.Float},

			// Normal field: department is a String
			"department": &gql.Field{
				Type: gql.String,
			},

			// Object field: departmentDetails is a Department
			"departmentDetails": &gql.Field{
				Type: departmentType,
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					switch employee := p.Source.(type) {
					case model.Employee:
						return map[string]interface{}{
							"name": employee.Department,
						}, nil
					case *model.Employee:
						if employee == nil {
							return nil, nil
						}
						return map[string]interface{}{
							"name": employee.Department,
						}, nil
					default:
						return nil, fmt.Errorf("invalid employee data")
					}
				},
			},
		},
	})

	// Input Type for create and update
	employeeInput := gql.NewInputObject(gql.InputObjectConfig{
		Name: "EmployeeInput",
		Fields: gql.InputObjectConfigFieldMap{
			"name":       &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
			"email":      &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
			"department": &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
			"salary":     &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.Float)},
		},
	})

	// Queries
	queryType := gql.NewObject(gql.ObjectConfig{
		Name: "Query",
		Fields: gql.Fields{
			"employees": &gql.Field{
				Type: gql.NewList(employeeType),
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					return server.service.GetAll()
				},
			},
			"employee": &gql.Field{
				Type: employeeType,
				Args: gql.FieldConfigArgument{
					"id": &gql.ArgumentConfig{
						Type: gql.NewNonNull(gql.ID),
					},
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

	// Mutations
	mutationType := gql.NewObject(gql.ObjectConfig{
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
					"id": &gql.ArgumentConfig{
						Type: gql.NewNonNull(gql.ID),
					},
					"input": &gql.ArgumentConfig{
						Type: gql.NewNonNull(employeeInput),
					},
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
				Type: gql.Boolean,
				Args: gql.FieldConfigArgument{
					"id": &gql.ArgumentConfig{
						Type: gql.NewNonNull(gql.ID),
					},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					id, err := parseGraphQLID(p.Args["id"])
					if err != nil {
						return false, err
					}

					if err := server.service.Delete(id); err != nil {
						return false, err
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

func (s *Server) Schema() *gql.Schema {
	return &s.schema
}

func parseGraphQLID(value interface{}) (uint, error) {
	id, err := strconv.ParseUint(fmt.Sprint(value), 10, 64)
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
	salary, okSalary := input["salary"].(float64)

	if !okName || !okEmail || !okDepartment || !okSalary {
		return model.Employee{}, fmt.Errorf("invalid employee input")
	}

	return model.Employee{
		Name:       name,
		Email:      email,
		Department: department,
		Salary:     salary,
	}, nil
}
