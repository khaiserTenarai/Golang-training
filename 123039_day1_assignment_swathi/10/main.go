
package main

import (
	"fmt"
	"os"
)

func main() {

	applicationName := os.Getenv("APP_NAME")
	port := os.Getenv("APP_PORT")

	if applicationName == "" {
		applicationName = "My Go Application"
	}

	if port == "" {
		port = "8080"
	}

	fmt.Println("===== APPLICATION DETAILS =====")
	fmt.Println("Application Name:", applicationName)
	fmt.Println("Port:", port)
}


