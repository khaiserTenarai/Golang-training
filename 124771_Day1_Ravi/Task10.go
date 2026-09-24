package main

import (
	"fmt"
	"os"
)

func main() {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "MyGoApp"
	}

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	fmt.Printf("Starting %s on port %s...\n", appName, appPort)
}
