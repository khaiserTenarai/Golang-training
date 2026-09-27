
// Day 4, Q9. Demonstrate pointer-based modification.

package main

import "fmt"

type Employee struct {
	Name   string
	Salary float64
}

// giveRaise takes a pointer, so changes made here are visible to the
// caller too, it's editing the original struct, not a copy of it.
func giveRaise(e *Employee, amount float64) {
	e.Salary += amount
}

func main() {
	gokul := Employee{Name: "Gokul", Salary: 60000}

	fmt.Println("Before raise:", gokul.Salary)
	giveRaise(&gokul, 5000)
	fmt.Println("After raise :", gokul.Salary)
}
