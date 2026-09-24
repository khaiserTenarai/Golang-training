// Package employee contains the Employee model and an in-memory Store.
package employee

import (
	"errors"
	"strings"
)

// Sentinel errors returned by the Store.
var (
	ErrNotFound     = errors.New("employee not found")
	ErrInvalidName  = errors.New("name must not be empty")
	ErrInvalidDept  = errors.New("department must not be empty")
	ErrInvalidSalry = errors.New("salary must be greater than zero")
)

// Employee is a single employee record.
type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
}

// Validate checks that all fields hold acceptable values.
func (e Employee) Validate() error {
	switch {
	case strings.TrimSpace(e.Name) == "":
		return ErrInvalidName
	case strings.TrimSpace(e.Department) == "":
		return ErrInvalidDept
	case e.Salary <= 0:
		return ErrInvalidSalry
	}
	return nil
}
