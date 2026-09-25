package main

import (
	"errors"
	"fmt"
)

var ErrInvalidAge = errors.New("age must be between 18 and 100")

func validateAge(age int) error {
	if age < 18 || age > 100 {
		return ErrInvalidAge
	}
	return nil
}

func main() {

	fmt.Println("\n***********************************")
	fmt.Println("11. Create custom validation errors.")
	fmt.Println("************************************")

	var age int

	fmt.Print("Enter Age: ")
	fmt.Scan(&age)

	err := validateAge(age)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Age is valid!")
	}
}