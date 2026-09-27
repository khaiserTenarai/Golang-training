
// Day 4, Q10. Demonstrate value vs pointer behavior.

package main

import "fmt"

type Employee struct {
	Name   string
	Salary float64
}


func modifyByValue(e Employee) {
	e.Salary = 99999
}

func modifyByPointer(e *Employee) {
	e.Salary = 99999
}

func main() {
	gokul := Employee{Name: "Gokul", Salary: 60000}

	modifyByValue(gokul)
	fmt.Println("After modifyByValue  :", gokul.Salary) // unchanged

	modifyByPointer(&gokul)
	fmt.Println("After modifyByPointer:", gokul.Salary) // changed
}
