package main

import (
	"fmt"
)

func check(age int) error {
	if age < 18 {
		return fmt.Errorf("age %d is not allowed; must be 18 or above", age)
	}

	return nil
}

func main() {
	age := 50

	err := check(age)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Age is valid")
}
