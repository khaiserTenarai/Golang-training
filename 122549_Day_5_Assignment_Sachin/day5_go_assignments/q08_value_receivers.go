// 8. Demonstrate value receivers.

package main

import "fmt"

type Employee struct {
	Name   string
	Salary float64
}

func (e Employee) Display() {
	fmt.Printf("Employee: %s, Salary: %.2f\n", e.Name, e.Salary)
}

func (e Employee) TrySalaryChange(newSalary float64) {
	e.Salary = newSalary
}

func main() {
	emp := Employee{Name: "sachin", Salary: 60000}

	emp.Display()

	emp.TrySalaryChange(80000)
	fmt.Println("After TrySalaryChange (original unchanged):", emp.Salary)
}
