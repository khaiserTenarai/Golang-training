package assignment_task9

import "errors"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func ValidateEmployee(emp Employee) error {
	if emp.ID <= 0 {
		return errors.New("invalid employee ID")
	}
	if emp.Name == "" {
		return errors.New("employee name cannot be empty")
	}
	if emp.Salary < 0 {
		return errors.New("salary cannot be negative")
	}
	return nil
}