package main

import (
	"fmt"
	"runtime"
	"os"
)

func main() {

	appName := "Employee Management"
	appVersion := "1.0.0"
	environment := os.Getenv("ENVIRONMENT")

	if environment == "" {
		environment = "development"
	}

	fmt.Println("Application Name:", appName)
	fmt.Println("Application Version:", appVersion)
	fmt.Println("Go Version:", runtime.Version())
	fmt.Println("Environment:", environment)
}