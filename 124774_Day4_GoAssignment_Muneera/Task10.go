package main

import "fmt"

func changeValue(salary float64) {
	salary = 40000
}
func changePointer(salary *float64) {
	*salary = 50000
}
func main() {
	salary := 30000.0
	fmt.Println("Original Salary:", salary)
	changeValue(salary)
	fmt.Println("After Value Function:", salary)

	changePointer(&salary)
	fmt.Println("After Pointer Function:", salary)

}
