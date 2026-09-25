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
	var a float64
	fmt.Println("Enter Salary=")
	fmt.Scan(&a)
	err := validateSalary(a)

	if err != nil {
		fmt.Println(err)
	}
}
