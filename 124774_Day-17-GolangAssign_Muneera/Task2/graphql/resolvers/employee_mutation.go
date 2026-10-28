package resolvers

import (
	"errors"
	"strings"

	"employee-management/model"
	"employee-management/service"
	gql "github.com/graphql-go/graphql"
)

func fromInput(input map[string]interface{}) model.Employee {
	return model.Employee{
		Name:       strings.TrimSpace(input["name"].(string)),
		Email:      strings.TrimSpace(input["email"].(string)),
		Department: strings.TrimSpace(input["department"].(string)),
		Salary:     input["salary"].(float64),
	}
}

func CreateEmployee(svc service.EmployeeService) gql.FieldResolveFn {
	return func(p gql.ResolveParams) (interface{}, error) {
		input := p.Args["input"].(map[string]interface{})
		employee := fromInput(input)
		if err := svc.Create(&employee); err != nil {
			return nil, err
		}
		return &employee, nil
	}
}

func UpdateEmployee(svc service.EmployeeService) gql.FieldResolveFn {
	return func(p gql.ResolveParams) (interface{}, error) {
		id := uint(p.Args["id"].(int))
		input := p.Args["input"].(map[string]interface{})
		employee := fromInput(input)
		return svc.Update(id, &employee)
	}
}

func DeleteEmployee(svc service.EmployeeService) gql.FieldResolveFn {
	return func(p gql.ResolveParams) (interface{}, error) {
		id := uint(p.Args["id"].(int))
		if err := svc.Delete(id); err != nil {
			return false, err
		}
		return true, nil
	}
}

// Kept as a helper for resolver-level not-found handling if needed.
var _ = errors.New
