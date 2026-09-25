package main

import (
	"fmt"
)

type ValidationError struct {
	Field   string
	Message string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("Validation failed on field '%s': %s", v.Field, v.Message)
}

func validateEmployee(name string, age int) error {
	if name == "" {
		return &ValidationError{
			Field:   "Name",
			Message: "cannot be empty",
		}
	}
	if age < 18 {
		return &ValidationError{
			Field:   "Age",
			Message: "must be 18 or older",
		}
	}
	return nil
}

func main() {
	err1 := validateEmployee("", 25)
	if err1 != nil {
		fmt.Println(err1)
	}

	err2 := validateEmployee("Alice", 16)
	if err2 != nil {
		fmt.Println(err2)
	}

	err3 := validateEmployee("Lakshmi", 24)
	if err3 == nil {
		fmt.Println("Lakshmi is a valid employee.")
	}
}