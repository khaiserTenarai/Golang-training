// 13. Demonstrate errors.Is.

package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("record not found")

func findItem(id int) error {
	if id != 1 {
		return fmt.Errorf("db lookup error: %w", ErrNotFound)
	}
	return nil
}

func main() {
	err := findItem(2)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			fmt.Println("Caught target error: ErrNotFound")
		} else {
			fmt.Println("Other error:", err)
		}
	}
}
