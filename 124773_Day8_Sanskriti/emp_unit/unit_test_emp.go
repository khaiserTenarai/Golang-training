package main

import "errors"

func ValidateName(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}

	return nil
}

func ValidateAge(age int) error {
	if age <= 0 {
		return errors.New("age must be greater than 0")
	}

	return nil
}

func ValidateSalary(salary float64) error {
	if salary <= 0 {
		return errors.New("salary must be greater than 0")
	}

	return nil
}