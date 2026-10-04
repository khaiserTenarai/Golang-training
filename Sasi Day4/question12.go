package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("record not found")

func getUser(id int) error {
	return fmt.Errorf("failed to fetch user %d: %w", id, ErrNotFound)
}

func main() {
	err := getUser(99)

	if err != nil {
		fmt.Println("Error:", err)

		if errors.Is(err, ErrNotFound) {
			fmt.Println("Page not Found")
		}
	}
}