package main

import "fmt"

const AppName = "Go Basics Demo"
const MaxUsers int = 100

func main(){

	fmt.Println("\n**************************************************************")
	fmt.Println("1. Create a program demonstrating Go variables and constants.")
	fmt.Println("****************************************************************")

	var username string = "Alice"
	var isLoggedIn bool
	age := 25
	salary := 75000.50

	fmt.Println("=== CONSTANTS ===")
	fmt.Println("Application Name:", AppName)
	fmt.Println("Maximum Users:   ", MaxUsers)

	fmt.Println("\n=== VARIABLES ===")
	fmt.Println("User: ", username, "Age: ", age, "Salary: ", salary)
	fmt.Println("Logged In:", isLoggedIn)

}