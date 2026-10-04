package main

import "fmt"

func main() {
	sales := [12]float32{
		11000,
		12000,
		13000,
		14000,
		13000,
		16000,
		15000,
		18000,
		17000,
		20000,
		19000,
		16000,
	}

	var total float32

	for _, sale := range sales {
		total = total + sale
	}

	average := total / 12

	fmt.Println("Total Sales:", total)
	fmt.Println("Average Sales:", average)
}