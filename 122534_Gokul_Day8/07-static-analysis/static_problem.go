package main

import "fmt"

func main() {
	name := "Ram"
	salary := 45000

	// BUG: %d is for integers, but we are passing 'name' which is a string!
	// The Go compiler will build this, but go vet will catch it.
	fmt.Printf("Employee Name: %d\n", name)
	fmt.Printf("Salary: %d\n", salary)
}
