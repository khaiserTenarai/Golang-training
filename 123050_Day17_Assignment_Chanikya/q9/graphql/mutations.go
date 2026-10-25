package graphql

import (
	"fmt"

	"employee-management/middleware"
	"employee-management/model"

	gql "github.com/graphql-go/graphql"
)

func newMutationType(
	server *Server,
	employeeType *gql.Object,
	inputType *gql.InputObject,
) *gql.Object {
	return gql.NewObject(gql.ObjectConfig{
		Name: "Mutation",
		Fields: gql.Fields{
			"createEmployee": &gql.Field{
				Type: employeeType,
				Args: gql.FieldConfigArgument{
					"input": &gql.ArgumentConfig{Type: gql.NewNonNull(inputType)},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					if middleware.RoleFromContext(p.Context) != "admin" {
						return nil, fmt.Errorf("forbidden: admin role required")
					}
					e, err := employeeInput(p.Args["input"])
					if err != nil {
						return nil, err
					}
					if err := server.service.Create(e); err != nil {
						return nil, err
					}
					return e, nil
				},
			},

			"updateEmployee": &gql.Field{
				Type: employeeType,
				Args: gql.FieldConfigArgument{
					"id":    &gql.ArgumentConfig{Type: gql.NewNonNull(gql.ID)},
					"input": &gql.ArgumentConfig{Type: gql.NewNonNull(inputType)},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					if middleware.RoleFromContext(p.Context) != "admin" {
						return nil, fmt.Errorf("forbidden: admin role required")
					}
					id, err := parseGraphQLID(p.Args["id"])
					if err != nil {
						return nil, err
					}
					e, err := employeeInput(p.Args["input"])
					if err != nil {
						return nil, err
					}
					return server.service.Update(id, e)
				},
			},

			"deleteEmployee": &gql.Field{
				Type: gql.NewNonNull(gql.Boolean),
				Args: gql.FieldConfigArgument{
					"id": &gql.ArgumentConfig{Type: gql.NewNonNull(gql.ID)},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					if middleware.RoleFromContext(p.Context) != "admin" {
						return nil, fmt.Errorf("forbidden: admin role required")
					}
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
}

func employeeInput(value interface{}) (*model.Employee, error) {
	input, ok := value.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid employee input")
	}

	name, okName := input["name"].(string)
	email, okEmail := input["email"].(string)
	department, okDepartment := input["department"].(string)
	salary, okSalary := input["salary"].(float64)

	if !okName || !okEmail || !okDepartment || !okSalary {
		return nil, fmt.Errorf("name, email, department and salary are required")
	}

	return &model.Employee{
		Name: name, Email: email, Department: department, Salary: salary,
	}, nil
}
