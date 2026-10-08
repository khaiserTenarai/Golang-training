package main

import (
	"fmt"
	"sync"
)

func calculateSalary(name string, salary float64) {
	fmt.Println("Calculating salary for ", name)
	bonus := salary * 0.10
	totalSalary := salary + bonus
	fmt.Println("Total salary of ", name, "is ", totalSalary)
}
func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		calculateSalary("Tom", 40000)
	}()
	go func() {
		defer wg.Done()
		calculateSalary("Max", 80000)
	}()
	go func() {
		defer wg.Done()
		calculateSalary("John", 60000)
	}()
	wg.Wait()
	fmt.Println("All salary calculations completed")
}
