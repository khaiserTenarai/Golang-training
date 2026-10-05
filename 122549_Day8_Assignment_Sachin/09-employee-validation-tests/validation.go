// 9. Write unit tests for employee validation.
//
// This file has the actual validation logic. validation_test.go (in this
// same folder) contains the tests for it. main.go is a small interactive
// program that reads real employee details from the user and calls these
// same functions - the tests don't replace that, they just check the
// logic itself works correctly under many different inputs.

package main

import (
	"errors"
	"strings"
)

func validateName(name string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name cannot be empty")
	}
	return nil
}

func validateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return errors.New("email must contain @")
	}
	return nil
}

func validateAge(age int) error {
	if age < 18 {
		return errors.New("age must be 18 or above")
	}
	return nil
}
