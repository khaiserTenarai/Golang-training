// 8. Demonstrate value receivers.

package main

import "fmt"

type Employee struct {
	Name   string
	Salary float64
}

func (e Employee) TryToGiveRaise(amount float64) {
	e.Salary = e.Salary + amount
	fmt.Println("Inside the method, salary is now:", e.Salary)
}

func main() {
	emp := Employee{Name: "Anita Sharma", Salary: 45000}
	fmt.Println("Before calling method:", emp.Salary)

	emp.TryToGiveRaise(5000)

	fmt.Println("After calling method:", emp.Salary)
	
}
