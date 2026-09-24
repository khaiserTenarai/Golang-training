package main
import "fmt"
func main() {
	sales := [12]float64{
		1000, 1200, 1500, 1100,
		1300, 1434, 1600, 1800,
		1706, 1901, 2005, 2200,
	}
	total := 0.0

	for i := 0; i < 12; i++ {
		total += sales[i]
	}

	average := total / 12

	fmt.Println("Total Sales:", total)
	fmt.Println("Average Sales:", average)
}
