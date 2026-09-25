
package main

import "fmt"

func main() {

	// Array containing sales for 12 months
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
		20000,
		21000,
		22000,
	}

	var total float64

	// Calculate total sales
	for i := 0; i < len(sales); i++ {
		total = total + sales[i]
	}

	// Calculate average sales
	average := total / float64(len(sales))

	fmt.Println("===== MONTHLY SALES =====")

	for i := 0; i < len(sales); i++ {
		fmt.Println("Month", i+1, ":", sales[i])
	}

	fmt.Println("\n===== SALES SUMMARY =====")
	fmt.Println("Total Sales:", total)
	fmt.Println("Average Sales:", average)
}


