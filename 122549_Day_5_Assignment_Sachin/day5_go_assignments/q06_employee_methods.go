// 6. Create Employee methods.

package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func (e Employee) GetDetails() string {
	return fmt.Sprintf("ID: %d | Name: %s | Salary: $%.2f", e.ID, e.Name, e.Salary)
}

func (e Employee) AnnualBonus(percent float64) float64 {
	return e.Salary * (percent / 100)
}

func main() {
	emp := Employee{ID: 10, Name: "sachin", Salary: 4500}

	fmt.Println(emp.GetDetails())
	fmt.Printf("Bonus (10%%): $%.2f\n", emp.AnnualBonus(10))
}
