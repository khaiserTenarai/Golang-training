package utility

import (
	"employee-management/model"
	"errors"
	"strings"
)

func ValidateEmp(employee model.Employee) error {
	if strings.TrimSpace(employee.Name) == "" {
		return errors.New("Name can't be empty")
	}
	if strings.TrimSpace(employee.Email) == "" {
		return errors.New("Email can't be empty")
	}
	if employee.Age < 18 {
		return errors.New("Employee age must be greater than 18")
	}
	return nil
}
