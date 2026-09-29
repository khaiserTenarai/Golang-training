package service

import (
	"errors"
	"fmt"

	"employee-management/model"
	"employee-management/utility"
)

var employees []model.Employee

func AddEmployee(employee model.Employee) error {

	employee.Name = utility.CleanString(employee.Name)
	employee.Email = utility.CleanString(employee.Email)

	// Validate name
	err := utility.ValidateName(employee.Name)

	if err != nil {
		return err
	}

	// Validate email
	err = utility.ValidateEmail(employee.Email)

	if err != nil {
		return err
	}

	// Validate age
	err = utility.ValidateAge(employee.Age)

	if err != nil {
		return err
	}

	// Validate salary
	err = utility.ValidateSalary(employee.Salary)

	if err != nil {
		return err
	}

	// Check duplicate employee
	for _, emp := range employees {

		if emp.ID == employee.ID {
			return utility.ErrDuplicateEmployee
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

	return model.Employee{}, fmt.Errorf(
		"cannot get employee with ID %d: %w",
		id,
		utility.ErrEmployeeNotFound,
	)
}

func DeleteEmployee(id int) error {

	for i, employee := range employees {

		if employee.ID == id {

			employees = append(
				employees[:i],
				employees[i+1:]...,
			)

			return nil
		}
	}

	return fmt.Errorf(
		"cannot delete employee with ID %d: %w",
		id,
		utility.ErrEmployeeNotFound,
	)
}

// Helper function to demonstrate errors.As
func IsValidationError(err error) bool {

	var validationError utility.ValidationError

	return errors.As(err, &validationError)
}