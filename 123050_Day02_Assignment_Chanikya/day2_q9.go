package main

import "fmt"

func main() {
	sales := [12]float64{
		12500.50, 15000.00, 13200.75, 18400.00,
		21000.25, 17500.00, 14200.50, 16800.00,
		19500.75, 22000.00, 25400.50, 28100.00,
	}

	var total float64 = 0

	for i := 0; i < len(sales); i++ {
		total = total + sales[i]
	}

	avg := total / 12

	fmt.Println("Total Annual Sales: ", total)
	fmt.Println("Average Monthly Sales: ", avg)
}
