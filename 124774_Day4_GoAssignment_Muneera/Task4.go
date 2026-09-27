package main

import "fmt"

func calculateSalary(sai ...float64) float64 {
	total := 0
	for i := 0; i < len(sal); i++ {
		total += sal[i]
	}
	return total
}

func main() {
	result, err := divide(10, 2)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
