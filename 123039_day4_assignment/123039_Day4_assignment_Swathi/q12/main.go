package main

import (
	"errors"
	"fmt"
)

var ErrInvalidSalary = errors.New("invalid salary")

func validateSalary(salary float64) error {

	if salary <= 0 {
		return fmt.Errorf("salary validation failed: %w", ErrInvalidSalary)
	}

	return nil
}

func main() {

	err := validateSalary(0)

	if err != nil {
		fmt.Println("Error:", err)
	}
}