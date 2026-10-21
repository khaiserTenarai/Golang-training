package main

import "fmt"

// Task 9: Input Struct Field Validation Layer
// Ensures invalid elements are intercepted before storage execution.
type UserInput struct {
	Email string
	Age   int
}

func (i *UserInput) Validate() error {
	if i.Email == "" {
		return fmt.Errorf("email address cannot be empty")
	}
	if i.Age < 18 {
		return fmt.Errorf("age restriction violation: must be 18+")
	}
	return nil
}

func main() {
	input := &UserInput{Email: "", Age: 16}
	err := input.Validate()
	fmt.Printf("[Task 9] Validation intercept result: %v\n", err)
}
