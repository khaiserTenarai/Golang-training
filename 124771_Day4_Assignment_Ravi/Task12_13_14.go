package main

import (
	"errors"
	"fmt"
)

var ErrInvalidAge = errors.New("invalid age")

// Custom error type
type ValidationError struct {
	Field string
	Value int
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s (value: %d)", e.Field, e.Msg, e.Value)
}

func checkAge(age int) error {
	if age < 18 {
		validationErr := &ValidationError{
			Field: "age",
			Value: age,
			Msg:   "must be 18 or above",
		}

		// Wrap both the sentinel error and custom error
		return fmt.Errorf("%w: %w", ErrInvalidAge, validationErr)
	}

	return nil
}

func main() {
	err := checkAge(17)

	if err != nil {
		fmt.Println("Error:", err)

		// Demonstrate errors.Is
		if errors.Is(err, ErrInvalidAge) {
			fmt.Println("errors.Is: This is an invalid age error")
		}

		// Demonstrate errors.As
		var validationErr *ValidationError

		if errors.As(err, &validationErr) {
			fmt.Println("errors.As:")
			fmt.Println("  Field:", validationErr.Field)
			fmt.Println("  Value:", validationErr.Value)
			fmt.Println("  Message:", validationErr.Msg)
		}

		return
	}

	fmt.Println("Age is valid")
}
