package main

import (
	"errors"
	"fmt"
)

var ErrInvalidSalary = errors.New("invalid salary")

func checkSalary(salary float64) error {
	if salary <= 0 {
		return ErrInvalidSalary
	}
	return nil
}

func validateEmployee(salary float64) error {
	err := checkSalary(salary)
	if err != nil {
		return fmt.Errorf("emplyee validation failed:%w", err)
	}
	return nil
}

func main() {
	err := validateEmployee(0)

	if errors.Is(err, ErrInvalidSalary) {
		fmt.Println("Error:Invalid Salary")
	} else {
		fmt.Println("Salary is valid")
	}
}
