// Day 4, Q14. Demonstrate errors.As.

package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Field   string
	Message string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", v.Field, v.Message)
}

func validateSalary(amount float64) error {
	if amount < 0 {
		return fmt.Errorf("validateSalary failed: %w", &ValidationError{
			Field:   "Salary",
			Message: "cannot be negative",
		})
	}
	return nil
}

func main() {
	err := validateSalary(-500)

	var valErr *ValidationError
	if errors.As(err, &valErr) {
		fmt.Println("Validation error while updating Gokul's salary:")
		fmt.Println("  Field  :", valErr.Field)
		fmt.Println("  Message:", valErr.Message)
	} else {
		fmt.Println("No validation error found")
	}
}
