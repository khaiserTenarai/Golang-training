package utility

import "errors"

func ValidateEmployee(name string, salary float64) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	if salary <= 0 {
		return errors.New("salary must be greater than zero")
	}
	return nil
}
