package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	message string
}

func (e ValidationError) Error() string {
	return e.message
}

func validateEmployee() error {
	return ValidationError{"Invalid employee salary"}
}

func main() {
	err := validateEmployee()

	var validationError ValidationError

	if errors.As(err, &validationError) {
		fmt.Println("Validation Error:", validationError.message)
	}
}