package resolvers

import (
	"errors"

	"employee-management/service"

	gql "github.com/graphql-go/graphql"
)

func EmployeeByID(svc service.EmployeeService) gql.FieldResolveFn {
	return func(p gql.ResolveParams) (interface{}, error) {
		id := uint(p.Args["id"].(int))

		employee, err := svc.GetByID(id)
		if err != nil {
			if service.IsNotFound(err) {
				return nil, errors.New("employee not found")
			}
			return nil, err
		}

		return employee, nil
	}
}

func Employees(svc service.EmployeeService) gql.FieldResolveFn {
	return func(p gql.ResolveParams) (interface{}, error) {
		return svc.GetAll()
	}
}
