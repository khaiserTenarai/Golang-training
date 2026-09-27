// 9. Demonstrate pointer-based modification.

package main

import "fmt"

type Employee struct {
	Name   string
	Salary float64
}

func giveRaise(emp *Employee, percentage float64) {
	emp.Salary += emp.Salary * (percentage / 100)
}

func main() {
	e := Employee{Name: "David", Salary: 50000}
	fmt.Println("Before:", e)

	giveRaise(&e, 10)
	fmt.Println("After raise:", e)
}
