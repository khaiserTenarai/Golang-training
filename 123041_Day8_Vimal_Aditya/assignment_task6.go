package main

import "fmt"

func main() {
	empID := 101
	name := "Vimal"

	fmt.Printf("Employee ID: %d, Name: %s\n", empID, name, "extra_argument")
}

// go vet assignment_task6.go