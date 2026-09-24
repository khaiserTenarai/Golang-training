// Package employee contains the Employee domain model and operations.
package employee

import (
	"errors"
	"fmt"

	"packagedesign/utils"
)

// Employee represents a company employee.
type Employee struct {
	ID     int
	Name   string
	Salary float64
}

// New validates input and creates an Employee with a normalised name.
func New(id int, name string, salary float64) (Employee, error) {
	if name == "" {
		return Employee{}, errors.New("name must not be empty")
	}
	if salary < 0 {
		return Employee{}, errors.New("salary must not be negative")
	}
	return Employee{ID: id, Name: utils.TitleCase(name), Salary: salary}, nil
}

// GiveRaise increases the salary by the given percentage.
func (e *Employee) GiveRaise(percent float64) {
	e.Salary += e.Salary * percent / 100
}

// String implements fmt.Stringer.
func (e Employee) String() string {
	return fmt.Sprintf("[%d] %s - %s", e.ID, e.Name, utils.FormatCurrency(e.Salary))
}
