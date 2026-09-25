package main

import (
	"fmt"
	"os"
)

func main() {
	// Default values
	appName := "CompanyManagerApp"
	appPort := "8080"

	if len(os.Args) > 1 {
		appName = os.Args[1]
	}

	if len(os.Args) > 2 {
		appPort = os.Args[2]
	}

	fmt.Println("--- App Setup ---")
	fmt.Println("App Name:", appName)
	fmt.Println("Port:", appPort)

	// use this to run or provide u'r own: go run assignment_task10.go EmployeeApp 9090
}