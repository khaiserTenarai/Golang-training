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
		13000,
		17000,
		18000,
		15000,
		19000,
		20000,
	}

	var total float64

	// Calculate total sales
	for _, sale := range sales {
		total += sale
	}

	// Calculate average
	average := total / float64(len(sales))

	fmt.Println("Total Sales:", total)
	fmt.Println("Average Monthly Sales:", average)
}