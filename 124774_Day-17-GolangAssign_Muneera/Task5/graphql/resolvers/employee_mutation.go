package resolvers

import (
	"employee-management/model"
	"employee-management/service"

	gql "github.com/graphql-go/graphql"
)

func fromInput(input map[string]interface{}) model.Employee {
	return model.Employee{
		Name:         input["name"].(string),
		Email:        input["email"].(string),
		DepartmentID: uint(input["departmentId"].(int)),
		Salary:       input["salary"].(float64),
	}
}

func CreateEmployee(svc service.EmployeeService) gql.FieldResolveFn {
	return func(p gql.ResolveParams) (interface{}, error) {
		input := p.Args["input"].(map[string]interface{})
		employee := fromInput(input)

		if err := svc.Create(&employee); err != nil {
			return nil, err
		}

		// Reload employee with its nested department.
		return svc.GetByID(employee.ID)
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
