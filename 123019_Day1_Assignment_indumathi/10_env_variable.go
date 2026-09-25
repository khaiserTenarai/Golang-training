package main

import (
	"fmt"
	"os"
)

func main() {
	appName := os.Getenv("my-app")
	if appName == "" {
		appName = "Myapplic"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8100"
	}

	fmt.Println("Starting Application:", appName)
	fmt.Println("Listening on Port:", port)
}
