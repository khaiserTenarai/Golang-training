package service

import (
	"employee-management/model"
	"employee-management/utility"
	"errors"
	"fmt"
)

var ErrEmployeeNotFound = errors.New("Employee not found")
var ErrDuplicateEmployee = errors.New("Employee already exists")

var employees []model.Employee

func AddEmployee(employee model.Employee) error {
	employee.Name = utility.CleanString(employee.Name)
	employee.Email = utility.CleanString(employee.Email)

	if err := utility.ValidateName(employee.Name); err != nil {
		return err
	}

	if err := utility.ValidateEmail(employee.Email); err != nil {
		return err
	}

	if err := utility.ValidateAge(employee.Age); err != nil {
		return err
	}

	if err := utility.ValidateSalary(employee.Salary); err != nil {
		return err
	}

	for _, emp := range employees {
		if emp.ID == employee.ID {
			return ErrDuplicateEmployee
		}
	}
	employees = append(employees, employee)
	return nil
}
func GetEmployee(id int) (model.Employee, error) {
	for _, employee := range employees {
		if employee.ID == id {
			return employee, nil
		}
	}
	return model.Employee{}, fmt.Errorf("Get employee failed: %w", ErrEmployeeNotFound)
}
func GetAllEmployees() []model.Employee {
	return employees
}
func UpdateEmployee(employee model.Employee) error {
	for i := range employees {
		if employees[i].ID == employee.ID {
			employee.Name = utility.CleanString(employee.Name)
			employee.Email = utility.CleanString(employee.Email)

			if err := utility.ValidateName(employee.Name); err != nil {
				return err
			}

			if err := utility.ValidateEmail(employee.Email); err != nil {
				return err
			}

			if err := utility.ValidateAge(employee.Age); err != nil {
				return err
			}

			if err := utility.ValidateSalary(employee.Salary); err != nil {
				return err
			}

			employees[i] = employee
			return nil
		}
	}
	return ErrEmployeeNotFound
}
func DeleteEmployee(id int) error {
	for i, employee := range employees {
		if employee.ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			return nil
		}
	}
	return ErrEmployeeNotFound
}
