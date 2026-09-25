package service

import (
	"errors"
	"fmt"

	"employee-management/model"
	"employee-management/utility"
)

var ErrEmployeeNotFound = errors.New("employee not found")

var ErrDuplicateEmployee = errors.New("employee already exists")

type EmployeeService struct {
	employees []model.Employee
}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{
		employees: []model.Employee{},
	}
}

func (service *EmployeeService) AddEmployee(employee model.Employee) error {

	employee.Name = utility.CleanString(employee.Name)
	employee.Email = utility.CleanString(employee.Email)

	if err := utility.ValidateName(employee.Name); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	if err := utility.ValidateEmail(employee.Email); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	if err := utility.ValidateAge(employee.Age); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	if err := utility.ValidateSalary(employee.Salary); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	for _, existingEmployee := range service.employees {

		if existingEmployee.ID == employee.ID {
			return fmt.Errorf(
				"cannot add employee %d: %w",
				employee.ID,
				ErrDuplicateEmployee,
			)
		}
	}

	service.employees = append(service.employees, employee)

	return nil
}

func (service *EmployeeService) GetEmployee(id int) (model.Employee, error) {

	for _, employee := range service.employees {

		if employee.ID == id {
			return employee, nil
		}
	}

	return model.Employee{}, fmt.Errorf(
		"employee ID %d: %w",
		id,
		ErrEmployeeNotFound,
	)
}

func (service *EmployeeService) DeleteEmployee(id int) error {

	for index, employee := range service.employees {

		if employee.ID == id {

			service.employees = append(
				service.employees[:index],
				service.employees[index+1:]...,
			)

			return nil
		}
	}

	return fmt.Errorf(
		"cannot delete employee %d: %w",
		id,
		ErrEmployeeNotFound,
	)
}
