package utility

import "errors"

var ErrInvalidName = errors.New("invalid employee name")
var ErrInvalidAge = errors.New("invalid employee age")
var ErrInvalidSalary = errors.New("invalid employee salary")
var ErrInvalidPosition = errors.New("invalid employee position")

var ErrEmployeeNotFound = errors.New("employee not found")
var ErrEmployeeExists = errors.New("employee already exists")
