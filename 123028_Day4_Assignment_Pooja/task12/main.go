package main

import (
	"errors"
	"fmt"
)

func findEmployee(id int) error {
	err := errors.New("employee not found")

	return fmt.Errorf("error while finding employee: %w", err)
}

func main() {
	err := findEmployee(101)

	fmt.Println(err)
}