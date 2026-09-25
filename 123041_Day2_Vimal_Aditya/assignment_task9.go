package main

import "fmt"

func main() {

	fmt.Println("\n***************************************************************************")
	fmt.Println("9. Create an array containing 12 months of sales and calculate total/average.")
	fmt.Println("*****************************************************************************")

	sales := [12]float64{
		10000, 12000, 15000, 11000, 13000, 16000,
		14000, 18000, 17000, 19000, 20000, 22000,
	}

	var total float64

	for i := 0; i < len(sales); i++ {
		total = total + sales[i]
	}

	average := total / float64(len(sales))

	fmt.Println("--- Sales Summary ---")
	fmt.Println("Total Sales   :", total)
	fmt.Println("Average Sales :", average)
}