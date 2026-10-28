package graphql

import (
	"errors"

	"day17-employee-platform/model"

	gql "github.com/graphql-go/graphql"
)

func buildMutation(s *Server, employeeType *gql.Object) *gql.Object {
	inputType := gql.NewInputObject(gql.InputObjectConfig{
		Name: "EmployeeInput",
		Fields: gql.InputObjectConfigFieldMap{
			"name": {
				Type: gql.NewNonNull(gql.String),
			},
			"email": {
				Type: gql.NewNonNull(gql.String),
			},
			"salary": {
				Type: gql.NewNonNull(gql.Float),
			},
			"departmentId": {
				Type: gql.NewNonNull(gql.Int),
			},
		},
	})

	return gql.NewObject(gql.ObjectConfig{
		Name: "Mutation",
		Fields: gql.Fields{
			"createEmployee": {
				Type: employeeType,
				Args: gql.FieldConfigArgument{
					"input": {Type: gql.NewNonNull(inputType)},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					if roleFromContext(p.Context) != "admin" {
						return nil, errors.New("admin access required")
					}

					input := p.Args["input"].(map[string]interface{})
					employee := inputToEmployee(input)

					if err := s.EmployeeService.Create(employee); err != nil {
						return nil, err
					}

					return employee, nil
				},
			},
			"updateEmployee": {
				Type: employeeType,
				Args: gql.FieldConfigArgument{
					"id":    {Type: gql.NewNonNull(gql.Int)},
					"input": {Type: gql.NewNonNull(inputType)},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					if roleFromContext(p.Context) != "admin" {
						return nil, errors.New("admin access required")
					}

					id := uint(p.Args["id"].(int))
					input := p.Args["input"].(map[string]interface{})
					employee := inputToEmployee(input)

					return s.EmployeeService.Update(id, employee)
				},
			},
			"deleteEmployee": {
				Type: gql.NewNonNull(gql.Boolean),
				Args: gql.FieldConfigArgument{
					"id": {Type: gql.NewNonNull(gql.Int)},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					if roleFromContext(p.Context) != "admin" {
						return nil, errors.New("admin access required")
					}

					id := uint(p.Args["id"].(int))
					if err := s.EmployeeService.Delete(id); err != nil {
						return nil, err
					}

					return true, nil
				},
			},
		},
	})
}

// Keep model imported clearly for beginners when reading the mutation package.
var _ *model.Employee
