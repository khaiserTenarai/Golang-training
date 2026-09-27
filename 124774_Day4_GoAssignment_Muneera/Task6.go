package main

import "fmt"

func employeeId() func() int {
	id := 100

	return func() int {
		id++
		return id
	}
}

func main() {
	getId := employeeId()

	fmt.Println("Employeee ID:", getId())
	fmt.Println("Employeee ID:", getId())
	fmt.Println("Employeee ID:", getId())
}
