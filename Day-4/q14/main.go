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
	return e.Field + ": " + e.Message
}

func validateEmployee(name string) error {

	if name == "" {
		return &ValidationError{
			Field:   "Name",
			Message: "name cannot be empty",
		}
	}

	return nil
}

func main() {

	err := validateEmployee("")

	var validationError *ValidationError

	if errors.As(err, &validationError) {
		fmt.Println("Field:", validationError.Field)
		fmt.Println("Message:", validationError.Message)
	}
}