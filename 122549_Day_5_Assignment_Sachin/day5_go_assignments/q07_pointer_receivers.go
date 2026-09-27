// 7. Demonstrate pointer receivers.

package main

import "fmt"

type Employee struct {
	Name   string
	Salary float64
}

func (e *Employee) UpdateSalary(newSalary float64) {
	e.Salary = newSalary
}

func (e *Employee) Promote(raise float64) {
	e.Salary += raise
}

func main() {
	emp := Employee{Name: "sachin", Salary: 50000}
	fmt.Println("Before update:", emp)

	emp.UpdateSalary(58000)
	fmt.Println("After UpdateSalary:", emp)

	emp.Promote(5000)
	fmt.Println("After Promote:", emp)
}
