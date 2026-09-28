package main

import (
	"errors"
	"fmt"
)

func main() {
	var age int
	fmt.Println("Enter age to check eligibility:")
	fmt.Scan(&age)
	err := errors.New("Invalid age")
	if age < 18 {
		err1 := fmt.Errorf("Age %d is not eligible: %w", age, err)
		fmt.Println(err1)
	} else {
		fmt.Println("Eligible")
	}
}
