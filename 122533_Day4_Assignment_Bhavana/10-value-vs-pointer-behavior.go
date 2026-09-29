// 10. Demonstrate value vs pointer behavior.

package main

import "fmt"

type Employee struct {
	Name   string
	Salary int
}

func raiseSalaryByValue(emp Employee) {
	emp.Salary = emp.Salary + 5000
}

func raiseSalaryByPointer(emp *Employee) {
	emp.Salary = emp.Salary + 5000
}

func main() {
	emp1 := Employee{Name: "Ravi", Salary: 30000}
	raiseSalaryByValue(emp1)
	fmt.Println("After raise by value:", emp1)

	emp2 := Employee{Name: "Anita", Salary: 30000}
	raiseSalaryByPointer(&emp2)
	fmt.Println("After raise by pointer:", emp2)
}
