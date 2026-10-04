package main

import (
	"fmt"
	"sync"
	"time"
)

func calculateBonus(empID int, salary float64, wg *sync.WaitGroup) {
	defer wg.Done()
	
	// Simulate work
	time.Sleep(50 * time.Millisecond)
	bonus := salary * 0.10
	fmt.Printf("Employee %d: Bonus calculated: $%.2f\n", empID, bonus)
}

func main() {
	var wg sync.WaitGroup

	salaries := []float64{50000, 65000, 120000, 80000}

	for i, salary := range salaries {
		wg.Add(1)
		go calculateBonus(i+1, salary, &wg)
	}

	wg.Wait()
	fmt.Println("All employee calculations completed.")
}