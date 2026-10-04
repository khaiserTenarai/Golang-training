package service

import (
	"errors"
	"fmt"

	"employee-management/model"
	"employee-management/utility"
)


var ErrEmployeeNotFound = errors.New("employee not found")
var ErrDuplicateEmployee = errors.New("employee already exists")


var employees []model.Employee


func AddEmployee(emp model.Employee) error {
	emp.Name = utility.CleanString(emp.Name)
	emp.Email = utility.CleanString(emp.Email)

	if err := utility.ValidateName(emp.Name); err != nil {
		return fmt.Errorf("add employee failed: %w", err)
	}
	if err := utility.ValidateEmail(emp.Email); err != nil {
		return fmt.Errorf("add employee failed: %w", err)
	}
	if err := utility.ValidateAge(emp.Age); err != nil {
		return fmt.Errorf("add employee failed: %w", err)
	}
	if err := utility.ValidateSalary(emp.Salary); err != nil {
		return fmt.Errorf("add employee failed: %w", err)
	}

	for _, existing := range employees {
		if existing.ID == emp.ID {
			return fmt.Errorf("add employee failed: %w", ErrDuplicateEmployee)
		}
	}

	employees = append(employees, emp)
	return nil
}

func GetEmployee(id int) (model.Employee, error) {
	for _, emp := range employees {
		if emp.ID == id {
			return emp, nil
		}
	}
	return model.Employee{}, fmt.Errorf("get employee failed: %w", ErrEmployeeNotFound)
}

func DeleteEmployee(id int) error {
	for i, emp := range employees {
		if emp.ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("delete employee failed: %w", ErrEmployeeNotFound)
}
