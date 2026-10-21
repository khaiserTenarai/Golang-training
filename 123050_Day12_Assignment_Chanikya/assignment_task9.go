package main

import (
	"errors"
	"fmt"
)

// Task 9: Input struct field validation layer
// Struct validating parameters manually before triggering downstream domain structures.

type UserInput struct {
	Email string
	Age   int
}

func (u *UserInput) Validate() error {
	if u.Email == "" {
		return errors.New("email domain parameter cannot be blank")
	}
	if u.Age < 18 {
		return errors.New("age constraint validation failed: must be 18+")
	}
	return nil
}

func main() {
	input := UserInput{Email: "", Age: 12}
	err := input.Validate()
	fmt.Println("Task 9 - Validation outcome:", err)
}
