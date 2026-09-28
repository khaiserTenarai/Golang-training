package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("Invalid age")
	err1 := fmt.Errorf("Validation failed: %w", err)
	fmt.Println("Error:", err1)
	if errors.Is(err1, err) {
		fmt.Println("The original error is present")
	}
}
