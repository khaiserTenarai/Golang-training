// 10. Demonstrate value vs pointer behavior.

package main

import "fmt"

type Employee struct {
	Name string
}

func updateNameByValue(e Employee, newName string) {
	e.Name = newName
}

func updateNameByPointer(e *Employee, newName string) {
	e.Name = newName
}

func main() {
	emp := Employee{Name: "Initial Name"}

	updateNameByValue(emp, "Value Name")
	fmt.Println("After value update:", emp.Name)

	updateNameByPointer(&emp, "Pointer Name")
	fmt.Println("After pointer update:", emp.Name)
}
