// 10. Environment Variables
//
// Read application name and port from environment variables.
// Provide default values.

package main

import (
	"fmt"
	"os"
)

func main() {
	user := os.Getenv("USER")

	if user == "" {
		user = "Guest"
	}

	fmt.Println("Welcome,", user)
}
