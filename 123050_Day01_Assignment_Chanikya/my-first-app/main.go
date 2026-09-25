package main

import (
	"fmt"

	// Import the UUID package we just downloaded
	"github.com/google/uuid"
)

func main() {
	// Generate a brand new, random UUID (Version 4)
	newID := uuid.New()

	fmt.Println("Generating a unique ID for a new user...")

	// Print the unique ID as a string
	fmt.Printf("Your new ID is: %s\n", newID.String())
}
