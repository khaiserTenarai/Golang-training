package main

import (
	"errors"
	"fmt"
)

func checkSalary(salary float64) error {
	if salary <= 0 {
		return errors.New("invalid salary")
	}
	return nil
}

func validateEmployee(salary float64) error {
	err := checkSalary(salary)

	if err != nil {
		return fmt.Errorf("employee  validation failed:%w", err)
	}
	return nil
}
func main() {
	err := validateEmployee(0)

	if err != nil {
		fmt.Println("Error:", err)
	}
}
