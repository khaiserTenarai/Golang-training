package main

import (
	"fmt"
	"sync"
)

func calculateSalary(name string, salary float64) {
	fmt.Println("Calculating salary for", name)

	bonus := salary * 10 / 100
	totalSalary := salary + bonus

	fmt.Println(name, "Total Salary:", totalSalary)
}

func main() {

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		calculateSalary("Amaira", 30000)
	}()

	go func() {
		defer wg.Done()
		calculateSalary("Priya", 35000)
	}()

	go func() {
		defer wg.Done()
		calculateSalary("Rahul", 40000)
	}()

	wg.Wait()

	fmt.Println("All employee calculations completed")
}
