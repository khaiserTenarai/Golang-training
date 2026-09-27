package main

import (
	"errors"
	"fmt"
)

type SalaryError struct {
	Message string
}

func (e SalaryError) Error() string {
	return e.Message
}

func checkSalary(salary float64) error {
	if salary <= 0 {
		return SalaryError{
			Message: "salary must be greater than 0",
		}
	}
	return nil
}

func main() {
	err := checkSalary(0)

	var salaryError SalaryError
	if errors.As(err, &salaryError) {
		fmt.Println("Salary Error:", salaryError.Message)
	} else {
		fmt.Println("Salary is valid")
	}
}
