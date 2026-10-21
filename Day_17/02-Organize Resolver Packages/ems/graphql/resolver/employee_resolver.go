package resolver

import (
	"fmt"
	"strconv"

	"employee-management/model"
	"employee-management/service"
)

type EmployeeResolver struct {
	service service.EmployeeService
}

func NewEmployeeResolver(s service.EmployeeService) *EmployeeResolver {
	return &EmployeeResolver{
		service: s,
	}
}

func (r *EmployeeResolver) GetAll() (interface{}, error) {
	return r.service.GetAll()
}

func (r *EmployeeResolver) GetByID(value interface{}) (interface{}, error) {
	id, err := parseID(value)
	if err != nil {
		return nil, err
	}

	return r.service.GetByID(id)
}

func (r *EmployeeResolver) Create(value interface{}) (interface{}, error) {
	employee, err := employeeFromInput(value)
	if err != nil {
		return nil, err
	}

	if err := r.service.Create(&employee); err != nil {
		return nil, err
	}

	return &employee, nil
}

func (r *EmployeeResolver) Update(
	idValue interface{},
	input interface{},
) (interface{}, error) {
	id, err := parseID(idValue)
	if err != nil {
		return nil, err
	}

	employee, err := employeeFromInput(input)
	if err != nil {
		return nil, err
	}

	return r.service.Update(id, employee)
}

func (r *EmployeeResolver) Delete(value interface{}) (interface{}, error) {
	id, err := parseID(value)
	if err != nil {
		return nil, err
	}

	if err := r.service.Delete(id); err != nil {
		return nil, err
	}

	return true, nil
}

func parseID(value interface{}) (uint, error) {
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

	if err != nil || id == 0 || id > uint64(^uint(0)) {
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
		return model.Employee{}, fmt.Errorf(
			"name, email, and department are required",
		)
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
