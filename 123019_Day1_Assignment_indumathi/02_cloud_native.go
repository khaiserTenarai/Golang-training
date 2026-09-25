// Day 1 - Task 2: Hello Cloud-Native Go

package main

import (
	"fmt"
	"runtime"
)

func main() {
	appName := "Hello cloudnative"
	appVersion := "1.0.0"
	envName := "Development"

	goVersion := runtime.Version()

	fmt.Println("Application Name:", appName)
	fmt.Println("Application Version:", appVersion)
	fmt.Println("Go Version:", goVersion)
	fmt.Println("Environment Name:", envName)
}
//  Just used 
// The result after running laptop:
// Application Name:Hello cloudnative
// Application Version: 1.0.0
// Go Version: go1.27.1
// Environment Name: Development
