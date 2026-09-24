package main

import (
	"fmt"
	"os"
	"runtime"
)

// Injected via -ldflags at build time
var (
	appName    = "cloud-native-app"
	appVersion = "dev"
)

func main() {
	// Read runtime environment variable (defaults to "development")
	envName := os.Getenv("APP_ENV")
	if envName == "" {
		envName = "development"
	}

	// Fetch Go runtime version automatically
	goVersion := runtime.Version()

	// Display Application Details

	fmt.Printf(" Application Name : %s\n", appName)
	fmt.Printf(" Application Version: %s\n", appVersion)
	fmt.Printf(" Go Version         : %s\n", goVersion)
	fmt.Printf(" Environment Name   : %s\n", envName)

}
