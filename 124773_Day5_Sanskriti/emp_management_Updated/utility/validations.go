package utility

import (
	"errors"
	"strings"
)

func ValidateAge(age int) error {

	if age <= 0 {
		return errors.New("age must be greater than 0")
	}

	return nil
}

func ValidateEntry(value string) error {

	if strings.TrimSpace(value) == "" {
		return errors.New("entry cannot be empty")
	}

	return nil
}

func ValidateSalary(salary float64) error {

	if salary <= 1 {
		return errors.New("salary must be greater than 1")
	}

	return nil
}
