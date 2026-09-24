package main

import "fmt"

func main() {
	sales := [12]float64{
		10000,
		20000,
		30000,
		11000,
		13000,
		30000,
		14000,
		14000,
		16000,
		10000,
		20000,
		22000,
	}

	var total float64

	for _, sale := range sales {
		total = total + sale
	}

	average := total / 12

	fmt.Println("Total Sales:", total)
	fmt.Println("Average Sales:", average)
}