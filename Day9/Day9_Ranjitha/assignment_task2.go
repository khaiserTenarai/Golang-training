package main

import (
	"fmt"
	"sync"
)

type Employee struct {
	ID         int
	Name       string
	BaseSalary float64
}

type CalculationResult struct {
	EmpID     int
	Name      string
	Bonus     float64
	Tax       float64
	NetSalary float64
}

func calculatePayroll(emp Employee, results chan<- CalculationResult, wg *sync.WaitGroup) {
	defer wg.Done()

	bonus := emp.BaseSalary * 0.10 
	tax := emp.BaseSalary * 0.20
	netSalary := emp.BaseSalary + bonus - tax

	results <- CalculationResult{
		EmpID:     emp.ID,
		Name:      emp.Name,
		Bonus:     bonus,
		Tax:       tax,
		NetSalary: netSalary,
	}
}

func main() {
	employees := []Employee{
		{ID: 101, Name: "Vimal", BaseSalary: 50000},
		{ID: 102, Name: "Aditya", BaseSalary: 60000},
		{ID: 103, Name: "Adi", BaseSalary: 75000},
		{ID: 104, Name: "Alex", BaseSalary: 45000},
	}

	var wg sync.WaitGroup
	results := make(chan CalculationResult, len(employees))

	for _, emp := range employees {
		wg.Add(1)
		go calculatePayroll(emp, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Printf("%-5s %-10s %-10s %-10s %-10s\n", "ID", "NAME", "BONUS", "TAX", "NET SALARY")
	fmt.Println("--------------------------------------------------")
	for res := range results {
		fmt.Printf("%-5d %-10s %-10.2f %-10.2f %-10.2f\n",
			res.EmpID, res.Name, res.Bonus, res.Tax, res.NetSalary)
	}
}