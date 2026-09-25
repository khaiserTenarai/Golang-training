package main

import (
	"fmt"
	"runtime"
)

const (
	appName    = "Cloud-Native Go"
	appVersion = "1.0.0"
	envName    = "development"
)

func main() {
	fmt.Println("Application Name:", appName)
	fmt.Println("Application Version:", appVersion)
	fmt.Println("Go Version:", runtime.Version())
	fmt.Println("Environment Name:", envName)
}
