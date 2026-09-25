package main

import (
	"fmt"
	"os"
)

func main() {
	appName := os.Getenv("Application_Name")
	port := os.Getenv("Port")
	if appName == "" {
		appName = "Employee App"
	}
	if port == "" {
		port = "8080"
	}
	fmt.Println("Application Name: ", appName)
	fmt.Println("Port: ", port)
}
