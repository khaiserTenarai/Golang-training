// 10. Implement composition instead of inheritance.

package main

import "fmt"


type Employee struct {
	Name   string
	Salary float64
}

func (e Employee) Describe() string {
	return fmt.Sprintf("%s earns %.2f", e.Name, e.Salary)
}

type Manager struct {
	Employee     
	TeamSize int
}

func main() {
	mgr := Manager{
		Employee: Employee{Name: "Priya Nair", Salary: 80000},
		TeamSize: 5,
	}

	fmt.Println(mgr.Describe())
	fmt.Println("Team Size:", mgr.TeamSize)

	fmt.Println("Manager Name:", mgr.Name)
}
