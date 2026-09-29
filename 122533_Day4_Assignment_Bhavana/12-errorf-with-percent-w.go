// 12. Use fmt.Errorf with %w.

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
	fmt.Println(err)

	original := errors.Unwrap(err)
	fmt.Println("Original error:", original)
}
