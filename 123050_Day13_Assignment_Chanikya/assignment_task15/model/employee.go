package model

import "fmt"

type Employee struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Role     string  `json:"role"`
	Salary   float64 `json:"salary"`
}

func (e *Employee) Validate() error {
	if e.ID == "" {
		return fmt.Errorf("employee primary identifier ID required")
	}
	if e.Name == "" {
		return fmt.Errorf("employee description Name field parameter required")
	}
	if e.Role == "" {
		return fmt.Errorf("authorization system profile role constraint target required")
	}
	return nil
}
