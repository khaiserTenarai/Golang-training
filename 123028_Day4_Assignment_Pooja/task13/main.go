package main

import (
	"errors"
	"fmt"
)

var ErrEmployeeNotFound = errors.New("employee not found")

func findEmployee(id int) error {
	return fmt.Errorf("employee ID %d: %w", id, ErrEmployeeNotFound)
}

func main() {
	err := findEmployee(101)

	if errors.Is(err, ErrEmployeeNotFound) {
		fmt.Println("Employee was not found")
	}
}