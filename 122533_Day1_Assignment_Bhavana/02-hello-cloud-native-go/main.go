// Day 1 - Task 2: Hello Cloud-Native Go
//
// Prints the basic metadata a cloud-native service usually exposes
// somewhere (health/info endpoint, startup logs, etc).
package main

import (
	"fmt"
	"runtime"
)

const (
	appName    = "EmployeeApp"
	appVersion = "1.0.0"
	envName    = "development"
)

func main() {
	fmt.Println("Application Name    :", appName)
	fmt.Println("Application Version :", appVersion)
	fmt.Println("Go Version          :", runtime.Version())
	fmt.Println("Environment         :", envName)
}
