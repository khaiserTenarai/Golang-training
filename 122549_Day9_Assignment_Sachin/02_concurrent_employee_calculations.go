package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func calculateBonus(salary float64, index int, results []float64, wg *sync.WaitGroup) {
	defer wg.Done()
	results[index] = salary * 0.10
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("How many employees? ")
	countText, _ := reader.ReadString('\n')
	count, err := strconv.Atoi(strings.TrimSpace(countText))
	if err != nil || count <= 0 {
		fmt.Println("Please enter a positive whole number.")
		return
	}

	salaries := make([]float64, count)
	for i := 0; i < count; i++ {
		fmt.Printf("Enter salary for employee %d: ", i+1)
		salaryText, _ := reader.ReadString('\n')
		salary, err := strconv.ParseFloat(strings.TrimSpace(salaryText), 64)
		if err != nil {
			fmt.Println("That wasn't a number, using 0.")
			salary = 0
		}
		salaries[i] = salary
	}

	bonuses := make([]float64, count)
	var wg sync.WaitGroup
	for i, salary := range salaries {
		wg.Add(1)
		go calculateBonus(salary, i, bonuses, &wg)
	}
	wg.Wait()

	fmt.Println("\nBonus results:")
	for i, bonus := range bonuses {
		fmt.Printf("Employee %d - salary %.2f -> bonus %.2f\n", i+1, salaries[i], bonus)
	}
}
