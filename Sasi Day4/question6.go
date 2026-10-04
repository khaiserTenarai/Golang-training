package main

import "fmt"

func employeeIDGenerator() func() int {
	id := 0
	return func() int {
		id++
		return id
	}
}

func main() {
	generateID := employeeIDGenerator()

	fmt.Println(generateID())
	fmt.Println(generateID())
	fmt.Println(generateID())

	generateAnotherDeptID := employeeIDGenerator()
	
	fmt.Println(generateAnotherDeptID())
	fmt.Println(generateAnotherDeptID())
}