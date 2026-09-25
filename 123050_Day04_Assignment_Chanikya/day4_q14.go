package main

import (
	"errors"
	"fmt"
)

type SalaryError struct {
	Salary float64
}

func (e *SalaryError) Error() string {
	return fmt.Sprintf("invalid salary: %.2f", e.Salary)
}

func validateSalary(salary float64) error {
	if salary <= 0 {
		return fmt.Errorf("salary validation failed: %w", &SalaryError{Salary: salary})
	}

	return nil
}

func main() {
	var a float64
	fmt.Println("Enter Salary=")
	fmt.Scan(&a)
	err := validateSalary(a)

	var salaryErr *SalaryError

	if errors.As(err, &salaryErr) {
		fmt.Println("Invalid salary:", salaryErr.Salary)
	}
}
