package resolvers

import (
	"errors"

	"employee-management/model"
	"employee-management/service"
	gql "github.com/graphql-go/graphql"
)

type EmployeeResolver struct {
	service service.EmployeeService
}

func NewEmployeeResolver(svc service.EmployeeService) *EmployeeResolver {
	return &EmployeeResolver{service: svc}
}

func (r *EmployeeResolver) EmployeeByID(p gql.ResolveParams) (interface{}, error) {
	id := uint(p.Args["id"].(int))
	employee, err := r.service.GetByID(id)
	if err != nil {
		if service.IsNotFound(err) {
			return nil, errors.New("employee not found")
		}
		return nil, err
	}
	return employee, nil
}

func (r *EmployeeResolver) Employees(p gql.ResolveParams) (interface{}, error) {
	return r.service.GetAll()
}

func (r *EmployeeResolver) CreateEmployee(p gql.ResolveParams) (interface{}, error) {
	input := p.Args["input"].(map[string]interface{})
	employee := fromInput(input)

	if err := r.service.Create(&employee); err != nil {
		return nil, err
	}
	return &employee, nil
}

func (r *EmployeeResolver) UpdateEmployee(p gql.ResolveParams) (interface{}, error) {
	id := uint(p.Args["id"].(int))
	input := p.Args["input"].(map[string]interface{})
	employee := fromInput(input)

	return r.service.Update(id, &employee)
}

func (r *EmployeeResolver) DeleteEmployee(p gql.ResolveParams) (interface{}, error) {
	id := uint(p.Args["id"].(int))
	if err := r.service.Delete(id); err != nil {
		return false, err
	}
	return true, nil
}

func fromInput(input map[string]interface{}) model.Employee {
	return model.Employee{
		Name:       input["name"].(string),
		Email:      input["email"].(string),
		Department: input["department"].(string),
		Salary:     input["salary"].(float64),
	}
}
