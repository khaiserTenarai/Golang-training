// 13. Demonstrate errors.Is.

package main

import (
	"errors"
	"fmt"
)

var errEmployeeNotFound = errors.New("employee not found")

func findEmployee(id int) error {
	if id != 1 {
		return fmt.Errorf("failed to load employee %d: %w", id, errEmployeeNotFound)
	}
	return nil
}

func main() {
	err := findEmployee(5)

	if errors.Is(err, errEmployeeNotFound) {
		fmt.Println("That employee does not exist.")
	} else {
		fmt.Println("Some other error happened:", err)
	}
}
