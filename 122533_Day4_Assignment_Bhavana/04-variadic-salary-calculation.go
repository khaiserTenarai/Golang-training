// 4. Create a variadic salary calculation function.

package main

import "fmt"


func totalSalary(salaries ...float64) float64 {
	var total float64
	for _, s := range salaries {
		total += s
	}
	return total
}

func main() {
	total := totalSalary(30000, 45000, 27000)
	fmt.Println("Total salary for 3 employees:", total)

	total = totalSalary(30000, 45000, 27000, 52000, 41000)
	fmt.Println("Total salary for 5 employees:", total)

	fmt.Println("Total salary for 0 employees:", totalSalary())
}
