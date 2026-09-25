package main

import "fmt"

//Hello Cloud-Native Go
func main() {
	applicationName := "Go Application"
	applicationVersion := "1.0.0"
	environmentName := "Development"
	goVersion := "go1.27.1"
	fmt.Println("Hello Cloud-Native Go")
	fmt.Println("Application Name: ", applicationName)
	fmt.Println("Application Version: ", applicationVersion)
	fmt.Println("Go Version: ", goVersion)
	fmt.Println("Environment Name: ", environmentName)
}
