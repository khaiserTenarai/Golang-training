// Day 2 - Task 9: Monthly Sales Array
package main

import "fmt"

func main() {
	sales := [12]float64{
		12000, 15500, 13200, 17800, 16000, 19500,
		21000, 20500, 18700, 22000, 24500, 26000,
	}

	var total float64
	for _, amount := range sales {
		total += amount
	}
	average := total / float64(len(sales))

	fmt.Println("Monthly Sales :", sales)
	fmt.Printf("Total Sales   : %.2f\n", total)
	fmt.Printf("Average Sales : %.2f\n", average)
}
