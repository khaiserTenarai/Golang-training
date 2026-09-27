

package main

import "fmt"

type ValidationError struct {
	Field   string
	Message string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on %s: %s", v.Field, v.Message)
}

func validateAge(age int) error {
	if age < 18 {
		return &ValidationError{Field: "Age", Message: "must be 18 or older"}
	}
	return nil
}

func main() {
	err := validateAge(16)
	if err != nil {
		fmt.Println(err)
	}

	err = validateAge(24)
	if err == nil {
		fmt.Println("Age 24 is valid for ranjitha")
	}
}
