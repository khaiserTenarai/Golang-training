// Task 9: An ARRAY of 12 months of sales: total, average and more.
// An array has a FIXED size that is part of its type: [12]float64.
package main

import (
	"fmt"
	"strings"
)

func main() {
	months := [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun",
		"Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

	// Monthly sales in lakh rupees
	sales := [12]float64{12.5, 15.0, 18.2, 14.8, 20.1, 22.4,
		19.6, 17.3, 21.0, 25.5, 28.9, 32.7}

	var total float64
	maxIdx, minIdx := 0, 0
	for i, amount := range sales {
		total += amount
		if amount > sales[maxIdx] {
			maxIdx = i
		}
		if amount < sales[minIdx] {
			minIdx = i
		}
	}
	average := total / float64(len(sales))

	fmt.Println("========= ANNUAL SALES REPORT (₹ lakh) =========")
	for i, amount := range sales {
		bar := strings.Repeat("█", int(amount))
		marker := ""
		if amount > average {
			marker = " ▲ above avg"
		}
		fmt.Printf("%s  %5.1f  %-33s%s\n", months[i], amount, bar, marker)
	}
	fmt.Println("-------------------------------------------------")
	fmt.Printf("Total sales    : ₹%.1f lakh\n", total)
	fmt.Printf("Average/month  : ₹%.2f lakh\n", average)
	fmt.Printf("Best month     : %s (₹%.1f lakh)\n", months[maxIdx], sales[maxIdx])
	fmt.Printf("Worst month    : %s (₹%.1f lakh)\n", months[minIdx], sales[minIdx])

	// Quarterly totals: slicing an array gives a slice view of it
	fmt.Println("\nQuarterly totals:")
	for q := 0; q < 4; q++ {
		quarter := sales[q*3 : q*3+3]
		sum := 0.0
		for _, v := range quarter {
			sum += v
		}
		fmt.Printf("  Q%d (%s-%s): ₹%.1f lakh\n", q+1, months[q*3], months[q*3+2], sum)
	}

	growth := (sales[11] - sales[0]) / sales[0] * 100
	fmt.Printf("\nGrowth Jan -> Dec: %.1f%%\n", growth)

	fmt.Println("\n=== Arrays are VALUES (copied on assignment) ===")
	backup := sales // full copy of all 12 values
	backup[0] = 999
	fmt.Println("backup[0] =", backup[0], "| sales[0] still =", sales[0])
	fmt.Println("len(sales) =", len(sales), "(fixed, can never grow)")
}
