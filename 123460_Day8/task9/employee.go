package employee

import (
	"errors"
	"strings"
)

type Employee struct {
	ID     int
	Name   string
	Age    int
	Email  string
	Salary float64
}

// Validate checks the structural integrity of an Employee record.
func (e Employee) Validate() error {
	if e.ID <= 0 {
		return errors.New("invalid employee ID: must be greater than 0")
	}
	if strings.TrimSpace(e.Name) == "" {
		return errors.New("invalid employee name: cannot be empty")
	}
	if e.Age < 18 || e.Age > 65 {
		return errors.New("invalid employee age: must be between 18 and 65")
	}
	if !strings.Contains(e.Email, "@") || !strings.Contains(e.Email, ".") {
		return errors.New("invalid employee email format")
	}
	if e.Salary <= 0 {
		return errors.New("invalid employee salary: must be positive")
	}
	return nil
}