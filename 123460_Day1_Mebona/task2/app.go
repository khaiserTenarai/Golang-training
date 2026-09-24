package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	
	appName := "Hello Cloud-Native Go"
	
	appVersion := "1.0.0"

	goVersion := runtime.Version()

	envName := os.Getenv("APP_ENV")
	if envName == "" {
		envName = "Development (default)"
	}


	fmt.Printf("Application Name    : %s\n", appName)
	fmt.Printf("Application Version : %s\n", appVersion)
	fmt.Printf("Go Version          : %s\n", goVersion)
	fmt.Printf("Environment Name    : %s\n", envName)
}