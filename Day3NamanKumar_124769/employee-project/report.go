package main

import "fmt"

// GenerateReport prints a summary report of all employees.
func GenerateReport() {
	total := 0.0
	for _, e := range Employees {
		total += e.Salary
	}
	fmt.Println("=== Employee Report ===")
	fmt.Printf("Headcount: %d\n", len(Employees))
	fmt.Printf("Total Salary: $%.2f\n", total)
	if len(Employees) > 0 {
		fmt.Printf("Average Salary: $%.2f\n", total/float64(len(Employees)))
	}
}

// GenerateRoleBreakdown prints headcount grouped by role.
func GenerateRoleBreakdown() {
	counts := map[string]int{}
	for _, e := range Employees {
		counts[e.Role]++
	}
	fmt.Println("=== Role Breakdown ===")
	for role, count := range counts {
		fmt.Printf("%s: %d\n", role, count)
	}
}
