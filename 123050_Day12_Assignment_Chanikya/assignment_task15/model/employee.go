package model

import "errors"

type Employee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

func (e *Employee) Validate() error {
	if e.Name == "" {
		return errors.New("employee name cannot be empty")
	}
	if e.Department == "" {
		return errors.New("employee department assignment is required")
	}
	if e.Salary <= 0 {
		return errors.New("salary context validation failed: must be positive numeric value")
	}
	return nil
}
