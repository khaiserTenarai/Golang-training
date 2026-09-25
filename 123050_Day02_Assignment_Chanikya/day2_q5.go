package main

import "fmt"

func main() {
	var basePay float64
	var allowance float64
	var taxRate float64
	fmt.Println("Enter base pay =")
	_, err := fmt.Scan(&basePay)
	if err != nil {
		fmt.Println("Invalid input.")
		return
	}
	fmt.Println("Enter allowance =")
	_, err = fmt.Scan(&allowance)
	if err != nil {
		fmt.Println("Invalid input.")
		return
	}
	fmt.Println("Enter taxRate =")
	_, err = fmt.Scan(&taxRate)
	if err != nil {
		fmt.Println("Invalid input.")
		return
	}

	fmt.Println("Base Pay:", basePay)
	fmt.Println("Allowance:", allowance)

	grossEarnings := basePay + allowance
	fmt.Println("Gross Earnings: ", grossEarnings)

	taxDeduction := grossEarnings * taxRate
	fmt.Println("Tax Deduction :", taxDeduction)

	netSalary := grossEarnings - taxDeduction

	fmt.Println("Final Net Salary: ", netSalary)

}
