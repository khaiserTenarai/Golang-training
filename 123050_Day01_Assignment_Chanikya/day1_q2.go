package main

import (
	"fmt"
	"runtime"
)

const (
	AppName    = "VS Code"
	AppVersion = "v1.4.2"
	EnvName    = "Production"
)

func main() {
	fmt.Println("Application Name:   ", AppName)
	fmt.Println("Application Version:", AppVersion)
	fmt.Println("Go Version:         ", runtime.Version())
	fmt.Println("Environment Name:   ", EnvName)
}
