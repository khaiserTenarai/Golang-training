package main

import "fmt"

func main() {

	sales := [12]int{
		1000,
		1200,
		1500,
		1100,
		1300,
		1600,
		1400,
		1700,
		1800,
		1100,
		1900,
		2000,
	}

	total := 0

	for i := 0; i < 12; i++ {
		total = total + sales[i]
	}

	average := total / 12

	fmt.Println("Total Sales:", total)
	fmt.Println("Average Sales:", average)
}
