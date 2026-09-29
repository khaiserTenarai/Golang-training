package model

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Email  string
	Salary float64
}

func (e Employee) Describe() string {
	return fmt.Sprintf("[%d] %s (%s) - Salary: %.2f", e.ID, e.Name, e.Email, e.Salary)
}
