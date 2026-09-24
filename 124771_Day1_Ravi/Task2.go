package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}
	AppName := "GoApp"
	AppVersion := "v1.0.0 "

	goVersion := runtime.Version()
	fmt.Printf(" Application Name : %s\n", AppName)
	fmt.Printf(" Application Version : %s\n", AppVersion)
	fmt.Printf(" Go Version : %s\n", goVersion)
	fmt.Printf(" Environment : %s\n", env)
}
