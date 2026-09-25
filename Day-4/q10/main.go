package main

import "fmt"

func changeValue(salary float64) {
	salary = 60000
}

func changePointer(salary *float64) {
	*salary = 60000
}

func main() {

	salary1 := 50000.0
	salary2 := 50000.0

	changeValue(salary1)
	changePointer(&salary2)

	fmt.Println("Value:", salary1)
	fmt.Println("Pointer:", salary2)
}