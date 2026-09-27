package utility

import "errors"

var (
	ErrInvalidName   = errors.New("invalid employee name")
	ErrInvalidAge    = errors.New("invalid employee age")
	ErrInvalidSalary = errors.New("invalid employee salary")
	ErrEmployeeNotFound = errors.New("employee not found")
)