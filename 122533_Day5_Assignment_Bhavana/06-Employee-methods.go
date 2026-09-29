// 6. Create Employee methods.

package main

import "fmt"

type Employee struct {
	Name   string
	Salary float64
}

func (e Employee) Describe() string {
	return fmt.Sprintf("%s earns %.2f", e.Name, e.Salary)
}

func (e Employee) AnnualSalary() float64 {
	return e.Salary * 12
}

func main() {
	emp := Employee{Name: "Anita Sharma", Salary: 45000}

	fmt.Println(emp.Describe())
	fmt.Println("Annual Salary:", emp.AnnualSalary())
}
