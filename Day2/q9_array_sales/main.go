// Q9. Create an array containing 12 months of sales and calculate

package main

import "fmt"

func main() {
	monthlySales := [12]float64{
		45000, 52000, 48000, 61000, 58000, 62000,
		70000, 65000, 59000, 72000, 68000, 80000,
	}

	total := 0.0
	for _, sale := range monthlySales {
		total += sale
	}
	average := total / float64(len(monthlySales))

	fmt.Println("Monthly Sales:", monthlySales)
	fmt.Printf("Total Sales: %.2f\n", total)
	fmt.Printf("Average Sales: %.2f\n", average)
}
