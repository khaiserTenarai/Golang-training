package main

import (
	"fmt"
	"sync"
	"time"
)

func calculateSalary(empID int, baseSalary float64, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(100 * time.Millisecond)
	netSalary := baseSalary * 1.10 // 10% bonus
	fmt.Printf("Employee %d: Net Salary = $%.2f\n", empID, netSalary)
}

func main() {
	var wg sync.WaitGroup
	employees := map[int]float64{101: 50000, 102: 60000, 103: 75000}

	for id, base := range employees {
		wg.Add(1)
		go calculateSalary(id, base, &wg)
	}

	wg.Wait()
}