package main

import "fmt"

type ValidationError struct {
	message string
}

func (e ValidationError) Error() string {
	return e.message
}

func validateSalary(salary int) error {
	if salary < 0 {
		return ValidationError{"Salary cannot be negative"}
	}

	return nil
}

func main() {
	err := validateSalary(-5000)

	if err != nil {
		fmt.Println("Error:", err)
	}
}