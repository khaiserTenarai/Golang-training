// 7. Demonstrate pointer receivers.

package main

import "fmt"

type Employee struct {
	Name   string
	Salary float64
}

func (e *Employee) GiveRaise(amount float64) {
	e.Salary = e.Salary + amount
}

func main() {
	emp := Employee{Name: "Anita Sharma", Salary: 45000}
	fmt.Println("Before raise:", emp.Salary)

	emp.GiveRaise(5000)
	fmt.Println("After raise:", emp.Salary)
}
