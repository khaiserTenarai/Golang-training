package main

import (
	"errors"
	"fmt"
)

func validateEmployee(name string, salary float64) error {

	if name == "" {
		return errors.New("employee name cannot be empty")
	}

	if salary <= 0 {
		return errors.New("salary must be greater than zero")
	}

	return nil
}

func main() {

	err := validateEmployee("Pallavi", 50000)

	if err != nil {
		fmt.Println("Validation Error:", err)
		return
	}

	fmt.Println("Employee is valid")
}