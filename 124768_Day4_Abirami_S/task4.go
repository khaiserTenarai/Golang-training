package main

import "fmt"

func calculateSalary(salaries ...int) int {
	total := 0
	for _, salary := range salaries {
		total += salary
	}
	return int(total)
}
func main() {
	basicPay := 50000
	hra := 5000
	bonus := 10000
	medicalAllowance := 2000
	totalSalary := calculateSalary(basicPay, hra, bonus, medicalAllowance)
	fmt.Println("Total Salary: ", totalSalary)
}
