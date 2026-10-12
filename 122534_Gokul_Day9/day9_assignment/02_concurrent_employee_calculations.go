package main

import (
	"fmt"
	"sync"
)

type salaryEmployee struct {
	ID     int
	Name   string
	Salary float64
}

func calculateNetSalary(e salaryEmployee, result *float64, wg *sync.WaitGroup) {
	defer wg.Done()
	bonus := e.Salary * 0.10
	tax := e.Salary * 0.05
	*result = e.Salary + bonus - tax
}

func main() {
	employees := []salaryEmployee{
		{1, "Asha", 30000},
		{2, "Ravi", 45000},
		{3, "Meena", 38000},
		{4, "Karthik", 52000},
	}

	netSalaries := make([]float64, len(employees))
	var wg sync.WaitGroup

	for i, emp := range employees {
		wg.Add(1)
		go calculateNetSalary(emp, &netSalaries[i], &wg)
	}
	wg.Wait()

	for i, emp := range employees {
		fmt.Printf("%d. %s -> net salary: %.2f\n", emp.ID, emp.Name, netSalaries[i])
	}
}
