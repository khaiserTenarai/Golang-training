package main

import "errors"

type Employee struct {
	Name   string
	Salary float64
}

func ValidateEmployee(employee Employee) error {
	if employee.Name == "" {
		return errors.New("employee name is required")
	}

	if employee.Salary <= 0 {
		return errors.New("salary must be greater than zero")
	}

	return nil
}
