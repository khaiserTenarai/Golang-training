package main

import (
	"errors"
	"fmt"
)

func validateSalary(salary float64) error {
	if salary <= 0 {
		return errors.New("salary must be greater than zero")
	}

	if salary < 10000 {
		return errors.New("salary must be at least 10000")
	}

	return nil
}

func main() {
	var a float64
	fmt.Println("Enter Salary=")
	fmt.Scan(&a)
	err := validateSalary(a)

	if err != nil {
		fmt.Println("Validation error:", err)
		return
	}

	fmt.Println("Salary is valid")
}
