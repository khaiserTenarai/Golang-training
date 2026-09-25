package main

import (
	"errors"
	"fmt"
)

type AgeError string

func (e AgeError) Error() string {
	return string(e)
}

func main() {

	fmt.Println("\n***********************")
	fmt.Println("14. Demonstrate errors.As")
	fmt.Println("*************************")

	err := fmt.Errorf("validation failed: %w", AgeError("age must be at least 18"))

	var targetErr AgeError

	if errors.As(err, &targetErr) {
		fmt.Println("Matched Custom Error")
		fmt.Println("Extracted Message:", targetErr)
	}
}