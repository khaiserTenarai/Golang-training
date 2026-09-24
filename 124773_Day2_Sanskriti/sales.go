package main

import "fmt"

func main() {
	sales := [12]float64{
		50000,
		55000,
		48000,
		60000,
		65000,
		70000,
		62000,
		68000,
		72000,
		75000,
		80000,
		85000,
	}

	months := [12]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	var total float64

	for i := 0; i < len(sales); i++ {
		fmt.Printf("%s: %.2f\n", months[i], sales[i])
		total += sales[i]
	}

	average := total / float64(len(sales))

	fmt.Printf("\nTotal Sales: %.2f\n", total)
	fmt.Printf("Average Sales: %.2f\n", average)
}