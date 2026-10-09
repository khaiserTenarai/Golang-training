package main

import (
	"fmt"
	"sync"
)

type Employee struct {
	Name   string
	Salary float64
	Bonus  float64
}

func calculateSalary(emp Employee, wg *sync.WaitGroup) {
	defer wg.Done()

	total := emp.Salary + emp.Bonus

	fmt.Printf("%s: Total Salary = %.2f\n", emp.Name, total)
}

func main() {
	employees := []Employee{
		{"Alice", 50000, 5000},
		{"Bob", 60000, 6000},
		{"Charlie", 55000, 5500},
		{"David", 70000, 7000},
	}

	var wg sync.WaitGroup

	for _, emp := range employees {
		wg.Add(1)
		go calculateSalary(emp, &wg)
	}

	wg.Wait()

	fmt.Println("All employee calculations completed.")
}
