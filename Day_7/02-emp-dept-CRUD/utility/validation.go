package utility

import (
	"errors"
	"strings"

	"department-management/model"
)

var (
	InvalidIDError           = errors.New("invalid ID")
	InvalidNameError         = errors.New("invalid name")
	InvalidAgeError          = errors.New("invalid age")
	InvalidEmailError        = errors.New("invalid email")
	InvalidDepartmentIDError = errors.New("invalid department ID")
)

func ValidateDepartment(
	department model.Department,
) error {

	if strings.TrimSpace(department.Name) == "" {
		return InvalidNameError
	}

	return nil
}

func ValidateEmployee(
	employee model.Employee,
) error {

	if strings.TrimSpace(employee.Name) == "" {
		return InvalidNameError
	}

	if employee.Age <= 18 || employee.Age >= 60 {
		return InvalidAgeError
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

func ValidateID(id int) error {

	if id <= 0 {
		return InvalidIDError
	}

	return nil
}

func ValidateDepartmentID(id int) error {

	if id <= 0 {
		return InvalidDepartmentIDError
	}

	return nil
}
