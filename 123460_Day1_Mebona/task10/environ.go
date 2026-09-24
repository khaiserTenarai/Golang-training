package main

import (
	"fmt"
	"os"
)

func getEnv(key string, defaultValue string) string {
	
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}

func main() {
	
	appName := getEnv("APP_NAME", "EmployeeService")
	port := getEnv("PORT", "8080")

	
	fmt.Printf("Starting application: %s\n", appName)
	fmt.Printf("Listening on port: %s\n", port)

	/*
	PS C:\Users\mebona.joseph\Desktop\day1_go> go run environ.go
	Starting application: EmployeeService
	Listening on port: 8080
	PS C:\Users\mebona.joseph\Desktop\day1_go> $env:APP_NAME="BillingService"
	PS C:\Users\mebona.joseph\Desktop\day1_go> $env:PORT="9090"
	PS C:\Users\mebona.joseph\Desktop\day1_go> go run environ.go             
	Starting application: BillingService
	Listening on port: 9090
	*/
}