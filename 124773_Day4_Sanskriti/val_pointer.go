
package main

import "fmt"

// Value
func changeValue(salary float64) {
	salary = 50000
}

// Pointer
func changePointer(salary *float64) {
	*salary = 50000
}

func main() {

	salary1 := 30000.0
	salary2 := 30000.0

	changeValue(salary1)
	changePointer(&salary2)

	fmt.Println("Value:", salary1)
	fmt.Println("Pointer:", salary2)
}