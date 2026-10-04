package main

import (
	"fmt"
	"sync"
)

type Employee struct {
	Name   string
	Salary float64
}

func calculateBonus(e Employee, results *[]string, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	bonus := e.Salary * 0.10
	mu.Lock()
	*results = append(*results, fmt.Sprintf("%s: bonus = %.2f", e.Name, bonus))
	mu.Unlock()
}

func main() {
	employees := []Employee{
		{"Anita", 50000},
		{"Rahul", 62000},
		{"Sara", 48000},
		{"Vikram", 71000},
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	results := []string{}

	for _, e := range employees {
		wg.Add(1)
		go calculateBonus(e, &results, &mu, &wg)
	}

	wg.Wait()

	for _, r := range results {
		fmt.Println(r)
	}
}
