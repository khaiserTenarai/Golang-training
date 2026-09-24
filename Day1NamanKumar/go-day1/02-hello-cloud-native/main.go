// Command hellocloud prints basic application metadata.
package main

import (
	"fmt"
	"os"
	"runtime"
)

const (
	appName    = "Hello Cloud-Native Go"
	appVersion = "1.0.0"
)

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	fmt.Println("==================================")
	fmt.Printf("Application : %s\n", appName)
	fmt.Printf("Version     : %s\n", appVersion)
	fmt.Printf("Go Version  : %s\n", runtime.Version())
	fmt.Printf("Environment : %s\n", env)
	fmt.Printf("Platform    : %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Println("==================================")
}
