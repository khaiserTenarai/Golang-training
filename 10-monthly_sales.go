package main

import "fmt"

func main() {

	// Sales for 12 months
	sales := [12]float64{
		10000,
		12000,
		15000,
		11000,
		14000,
		16000,
		18000,
		17000,
		19000,
		21000,
		22000,
		25000,
	}

	total := 0.0

	// Calculate total
	for _, sale := range sales {
		total += sale
	}

	// Calculate average
	average := total / float64(len(sales))

	fmt.Println("Total Sales:", total)
	fmt.Println("Average Monthly Sales:", average)
}
