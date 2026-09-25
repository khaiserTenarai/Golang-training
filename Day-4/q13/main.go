package main

import (
	"errors"
	"fmt"
)

var ErrEmployeeNotFound = errors.New("employee not found")

func findEmployee(id int) error {

	if id != 101 {
		return fmt.Errorf("search failed: %w", ErrEmployeeNotFound)
	}

	return nil
}

func main() {

	err := findEmployee(999)

	if errors.Is(err, ErrEmployeeNotFound) {
		fmt.Println("Employee not found")
	} else {
		fmt.Println("Employee found")
	}
}