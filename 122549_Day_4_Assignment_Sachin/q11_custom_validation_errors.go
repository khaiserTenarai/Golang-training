// 11. Create custom validation errors.

package main

import "fmt"

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on '%s': %s", e.Field, e.Message)
}

func validateAge(age int) error {
	if age < 18 {
		return &ValidationError{Field: "Age", Message: "must be at least 18"}
	}
	return nil
}

func main() {
	err := validateAge(16)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
