// 11. Create custom validation errors.

package main

import "fmt"


type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Message)
}

func validateAge(age int) error {
	if age < 18 {
		return &ValidationError{Field: "age", Message: "must be 18 or older"}
	}
	return nil
}

func main() {
	err := validateAge(15)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = validateAge(25)
	if err == nil {
		fmt.Println("Age 25 is valid")
	}
}
