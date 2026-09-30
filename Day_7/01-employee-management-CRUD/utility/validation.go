package utility

import (
	"errors"
	"strings"

	"ems/model"
)

var (
	InvalidEmployeeIDError   = errors.New("invalid employee ID")
	InvalidEmployeeNameError = errors.New("invalid employee name")
	InvalidEmployeeAgeError  = errors.New("invalid employee age")
	InvalidEmailError        = errors.New("invalid employee email")
)

func ValidateEmployee(employee model.Employee) error {

	if strings.TrimSpace(employee.Name) == "" {
		return InvalidEmployeeNameError
	}

	if employee.Age <= 18 || employee.Age >= 60 {
		return InvalidEmployeeAgeError
	}

	return nil
}

func ValidateEmployeeID(id int) error {

	if id <= 0 {
		return InvalidEmployeeIDError
	}

	return nil
}

func ValidateEmail(email string) error {

	email = strings.TrimSpace(email)

	if len(email) <= 3 || len(email) >= 100 {
		return InvalidEmailError
	}

	if !strings.Contains(email, "@") ||
		!strings.Contains(email, ".") {
		return InvalidEmailError
	}

	return nil
}