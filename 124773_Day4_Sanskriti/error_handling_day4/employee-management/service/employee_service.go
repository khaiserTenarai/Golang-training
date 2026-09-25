package service

import (
	"errors"
	"fmt"

	"employee-management/dao"
	"employee-management/model"
	"employee-management/utility"
)

func AddEmployee(employee model.Employee) error {

	// Clean strings
	employee.Name = utility.CleanString(employee.Name)
	employee.Email = utility.CleanString(employee.Email)

	// Validate name
	if err := utility.ValidateName(employee.Name); err != nil {
		return fmt.Errorf("employee validation failed: %w", err)
	}

	// Validate email
	if err := utility.ValidateEmail(employee.Email); err != nil {
		return fmt.Errorf("employee validation failed: %w", err)
	}

	// Validate age
	if err := utility.ValidateAge(employee.Age); err != nil {
		return fmt.Errorf("employee validation failed: %w", err)
	}

	// Validate salary
	if err := utility.ValidateSalary(employee.Salary); err != nil {
		return fmt.Errorf("employee validation failed: %w", err)
	}

	// Add to DAO
	err := dao.AddEmployee(employee)

	if err != nil {
		return fmt.Errorf("could not add employee: %w", err)
	}

	return nil
}

func GetEmployee(id int) (model.Employee, error) {

	employee, err := dao.GetEmployee(id)

	if err != nil {

		if errors.Is(err, utility.ErrEmployeeNotFound) {
			return model.Employee{}, fmt.Errorf(
				"get employee failed: %w",
				err,
			)
		}

		return model.Employee{}, err
	}

	return employee, nil
}

func DeleteEmployee(id int) error {

	err := dao.DeleteEmployee(id)

	if err != nil {

		if errors.Is(err, utility.ErrEmployeeNotFound) {
			return fmt.Errorf(
				"delete employee failed: %w",
				err,
			)
		}

		return err
	}

	return nil
}
