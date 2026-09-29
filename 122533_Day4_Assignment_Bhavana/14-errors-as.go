// 14. Demonstrate errors.As.

package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Message)
}

func validateAge(age int) error {
	if age < 18 {
		return fmt.Errorf("could not create employee: %w", &ValidationError{
			Field:   "age",
			Message: "must be 18 or older",
		})
	}
	return nil
}

func main() {
	err := validateAge(15)

	var validationErr *ValidationError
	if errors.As(err, &validationErr) {
		fmt.Println("Problem field:", validationErr.Field)
		fmt.Println("Message:", validationErr.Message)
	} else {
		fmt.Println("Some other error happened:", err)
	}
}
