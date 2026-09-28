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
func main() {
	var err error = ValidationError{"Age must be 18 or above"}
	if err != nil {
		var validationErr ValidationError
		if errors.As(err, &validationErr) {
			fmt.Println("Validation error: ", validationErr)
		}
	}
}
