// 2. Create a function returning multiple values.

package main

import "fmt"

func calculateStats(salaries []float64) (float64, float64) {
	if len(salaries) == 0 {
		return 0, 0
	}

	var total float64
	for _, s := range salaries {
		total += s
	}
	avg := total / float64(len(salaries))
	return total, avg
}

func main() {
	salaries := []float64{45000, 55000, 65000}
	total, average := calculateStats(salaries)

	fmt.Printf("Total: %.2f\n", total)
	fmt.Printf("Average: %.2f\n", average)
}
