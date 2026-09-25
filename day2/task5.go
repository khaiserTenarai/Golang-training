package main

import "fmt"

func main() {
	basicSalary := 50000.00
	hraAllowance := basicSalary * 0.20
	daAllowance := basicSalary * 0.10
	bonus := 5000.00
	taxDeduction := basicSalary * 0.12

	grossSalary := basicSalary + hraAllowance + daAllowance + bonus
	netSalary := grossSalary - taxDeduction

	fmt.Printf("Basic Salary: $%.2f\n", basicSalary)
	fmt.Printf("HRA (20%%):    $%.2f\n", hraAllowance)
	fmt.Printf("DA (10%%):     $%.2f\n", daAllowance)
	fmt.Printf("Bonus:        $%.2f\n", bonus)
	fmt.Printf("Gross Salary: $%.2f\n", grossSalary)
	fmt.Printf("Tax Deduct:  -$%.2f\n", taxDeduction)

	fmt.Printf("Net Salary:   $%.2f\n", netSalary)
}
