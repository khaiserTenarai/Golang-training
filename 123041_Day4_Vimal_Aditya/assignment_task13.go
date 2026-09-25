package main

import (
	"errors"
	"fmt"
)

var ErrInvalidAge = errors.New("age must be between 18 and 100")

func validateAge(age int) error {
	if age < 18 || age > 100 {
		return fmt.Errorf("user input error (age = %d): %w", age, ErrInvalidAge)
	}
	return nil
}

func main() {

	fmt.Println("\n***********************")
	fmt.Println("13. Demonstrate errors.Is")
	fmt.Println("*************************")

	var age int

	fmt.Print("Enter Age: ")
	fmt.Scan(&age)

	err := validateAge(age)

	if err != nil {
		fmt.Println("Full Error:", err)

		if errors.Is(err, ErrInvalidAge) {
			fmt.Println("Matched: cause is ErrInvalidAge.")
		}
	} else {
		fmt.Println("Age is valid!")
	}
}