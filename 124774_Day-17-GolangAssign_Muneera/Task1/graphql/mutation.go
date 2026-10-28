package graphql

import (
	"fmt"
	"strconv"

	gql "github.com/graphql-go/graphql"

	"employee-management/model"
	"employee-management/service"
)

func NewMutationType(svc service.EmployeeService) *gql.Object {
	return gql.NewObject(gql.ObjectConfig{
		Name: "Mutation",
		Fields: gql.Fields{
			"createEmployee": &gql.Field{
				Type: EmployeeType,
				Args: gql.FieldConfigArgument{
					"input": &gql.ArgumentConfig{
						Type: gql.NewNonNull(EmployeeInputType),
					},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					input := p.Args["input"].(map[string]interface{})

					employee := model.Employee{
						Name:       input["name"].(string),
						Email:      input["email"].(string),
						Department: input["department"].(string),
						Salary:     input["salary"].(float64),
					}

					if err := svc.Create(&employee); err != nil {
						return nil, err
					}

					return employee, nil
				},
			},

			"updateEmployee": &gql.Field{
				Type: EmployeeType,
				Args: gql.FieldConfigArgument{
					"id": &gql.ArgumentConfig{
						Type: gql.NewNonNull(gql.ID),
					},
					"input": &gql.ArgumentConfig{
						Type: gql.NewNonNull(EmployeeInputType),
					},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					idText := fmt.Sprint(p.Args["id"])
					idValue, err := strconv.ParseUint(idText, 10, 32)
					if err != nil || idValue == 0 {
						return nil, fmt.Errorf("invalid employee ID")
					}

					input := p.Args["input"].(map[string]interface{})
					employee := model.Employee{
						Name:       input["name"].(string),
						Email:      input["email"].(string),
						Department: input["department"].(string),
						Salary:     input["salary"].(float64),
					}

					return svc.Update(uint(idValue), employee)
				},
			},

			"deleteEmployee": &gql.Field{
				Type: gql.NewNonNull(gql.Boolean),
				Args: gql.FieldConfigArgument{
					"id": &gql.ArgumentConfig{
						Type: gql.NewNonNull(gql.ID),
					},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					idText := fmt.Sprint(p.Args["id"])
					idValue, err := strconv.ParseUint(idText, 10, 32)
					if err != nil || idValue == 0 {
						return false, fmt.Errorf("invalid employee ID")
					}

					if err := svc.Delete(uint(idValue)); err != nil {
						return false, err
					}

					return true, nil
				},
			},
		},
	})
}
