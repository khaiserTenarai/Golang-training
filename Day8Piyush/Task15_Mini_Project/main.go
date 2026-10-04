// Task 15: Day 8 Mini Project — Tested Employee Service
// A complete Employee Service demonstrating:
//   - Unit tests, table-driven tests, benchmarks
//   - Test coverage (go test -cover)
//   - Proper formatting (gofmt)
//   - Static analysis clean (go vet)
//   - Race-safe (go test -race)
//
// Run the demo:    go run main.go
// Run tests:       go test -v ./service/...
// Race detection:  go test -race ./service/...
// Coverage:        go test -cover ./service/...
// Benchmarks:      go test -bench=. -benchmem ./service/...
// Formatting:      gofmt -l .
// Vet:             go vet ./...

package main

import (
	"fmt"

	"employee_service/models"
	"employee_service/service"
)

func main() {
	fmt.Println("===== Day 8 Mini Project: Tested Employee Service =====\n")

	svc := service.NewEmployeeService()

	// Add employees
	emp1, _ := svc.Add(models.Employee{
		Name: "Piyush", Email: "piyush@test.com", Age: 25,
		Department: "Engineering", Salary: 60000,
	})
	fmt.Printf("Added: ID=%d, Name=%s\n", emp1.ID, emp1.Name)

	emp2, _ := svc.Add(models.Employee{
		Name: "Alice", Email: "alice@test.com", Age: 30,
		Department: "Design", Salary: 55000,
	})
	fmt.Printf("Added: ID=%d, Name=%s\n", emp2.ID, emp2.Name)

	// List all
	fmt.Printf("\nTotal employees: %d\n", svc.Count())
	for _, emp := range svc.GetAll() {
		fmt.Printf("  [%d] %s - %s - %.0f\n", emp.ID, emp.Name, emp.Department, emp.Salary)
	}

	// Salary calculations
	fmt.Println("\n--- Salary Calculations ---")

	net, _ := service.CalculateNetSalary(60000, 20)
	fmt.Printf("Net salary (60k, 20%% tax): %.0f\n", net)

	annual := service.CalculateAnnualSalary(60000)
	fmt.Printf("Annual salary: %.0f\n", annual)

	bonus, _ := service.CalculateBonus(60000, 4)
	fmt.Printf("Bonus (rating 4): %.0f\n", bonus)

	bracket := service.GetTaxBracket(annual)
	fmt.Printf("Tax bracket for %.0f: %s\n", annual, bracket)

	// Validation demo
	fmt.Println("\n--- Validation Demo ---")
	_, err := svc.Add(models.Employee{Name: "", Email: "bad", Age: 10, Department: "", Salary: -100})
	fmt.Printf("Invalid employee error: %v\n", err)

	_, err = svc.Add(models.Employee{Name: "Piyush", Email: "piyush@test.com", Age: 25, Department: "IT", Salary: 50000})
	fmt.Printf("Duplicate email error: %v\n", err)

	// Delete
	svc.Delete(2)
	fmt.Printf("\nAfter deleting ID 2, count: %d\n", svc.Count())

	fmt.Println("\n===== Mini Project Complete =====")
	fmt.Println("Run 'go test -v -cover -race ./service/...' to see full test suite.")
}
