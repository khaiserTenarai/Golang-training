package main

import (
	"errors"
	"fmt"
)

var InvalidAgeErr = errors.New("Not Eligible for voting")

func validateAge(age int) error {
	if age < 18 {
		return InvalidAgeErr
	}
	return nil
}
func main() {
	var age int
	fmt.Println("Enter age to check eligibility:")
	fmt.Scan(&age)
	err := validateAge(age)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Eligible for voting")
	}
}
