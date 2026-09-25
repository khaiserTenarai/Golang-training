package main

import "fmt"

func main() {
	
	monthlySales := [12]float64{
		12500.50, 15000.00, 13400.75, 18200.00,
		16500.25, 19000.00, 21000.50, 20500.00,
		17800.00, 22400.80, 25000.00, 28900.00,
	}

	totalSales := 0.0
	for _, sales := range monthlySales {
		totalSales += sales
	}

	averageSales := totalSales / float64(len(monthlySales))


	for month, sales := range monthlySales {
		fmt.Printf("Month %2d: $%.2f\n", month+1, sales)
	}
	
	fmt.Printf("Total Sales:   $%.2f\n", totalSales)
	fmt.Printf("Average Sales: $%.2f\n", averageSales)
}