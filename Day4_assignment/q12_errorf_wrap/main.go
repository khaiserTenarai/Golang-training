

package main

import (
	"errors"
	"fmt"
)

var ErrRecordNotFound = errors.New("record not found")

func findEmployee(id int) error {
	if id != 1 {
		return fmt.Errorf("findEmployee(%d): %w", id, ErrRecordNotFound)
	}
	return nil
}

func main() {
	err := findEmployee(99)
	if err != nil {
		fmt.Println("Got error:", err)
		fmt.Println("Is it ErrRecordNotFound?", errors.Is(err, ErrRecordNotFound))
	}

	err = findEmployee(1)
	fmt.Println("Lookup for ranjitha's record, error:", err)
}
