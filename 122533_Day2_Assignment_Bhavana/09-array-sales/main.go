// 9. Create an array containing 12 months of sales and calculate total/average.

package main

import "fmt"

func main() {
	
	monthlySales := [12]float64{
		12000, 15000, 11000, 17500, 20000, 21000,
		19500, 18000, 16000, 22000, 23500, 25000,
	}

	var total float64
	for _, sale := range monthlySales {
		total += sale
	}

	average := total / float64(len(monthlySales))

	fmt.Println("Monthly Sales:", monthlySales)
	fmt.Println("Total Sales for the year:", total)
	fmt.Printf("Average Monthly Sales: %.2f\n", average)
}
