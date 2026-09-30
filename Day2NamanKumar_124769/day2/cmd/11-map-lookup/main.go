// Task 11: Employee lookup using a MAP.
// A map stores key -> value pairs with near-instant O(1) lookup.
package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
}

func main() {
	// Create a map with a literal: key = employee ID
	byID := map[int]Employee{
		101: {101, "Anita Rao", "Engineering", 90000},
		102: {102, "Rahul Mehta", "HR", 45000},
		103: {103, "Priya Nair", "Finance", 60000},
		104: {104, "Arjun Singh", "Engineering", 70000},
	}

	// Add and update
	byID[105] = Employee{105, "Kavya Sharma", "Marketing", 55000}
	e := byID[102]
	e.Salary = 50000 // map values aren't addressable: copy, change, store back
	byID[102] = e

	// Lookup IDs from the command line, or use defaults
	ids := []string{"103", "999", "abc"}
	if len(os.Args) > 1 {
		ids = os.Args[1:]
	}
	fmt.Println("=== Lookup by ID (comma-ok idiom) ===")
	for _, s := range ids {
		id, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("%-5s -> invalid ID\n", s)
			continue
		}
		if emp, ok := byID[id]; ok {
			fmt.Printf("%-5d -> found: %s, %s, ₹%.0f\n", id, emp.Name, emp.Department, emp.Salary)
		} else {
			fmt.Printf("%-5d -> not found\n", id)
		}
	}

	fmt.Println("\nWithout comma-ok, a missing key silently returns the zero value:")
	fmt.Printf("byID[999] = %+v\n", byID[999])

	fmt.Println("\n=== Delete ===")
	delete(byID, 105)
	delete(byID, 999) // deleting a missing key is safe (no error)
	_, exists := byID[105]
	fmt.Println("105 exists after delete:", exists, "| total:", len(byID))

	fmt.Println("\n=== Iterate in sorted order (map order is random) ===")
	keys := make([]int, 0, len(byID))
	for id := range byID {
		keys = append(keys, id)
	}
	sort.Ints(keys)
	for _, id := range keys {
		fmt.Printf("%d  %s\n", id, byID[id].Name)
	}

	fmt.Println("\n=== Second index: name -> ID (lookup by name) ===")
	byName := make(map[string]int)
	for id, emp := range byID {
		byName[emp.Name] = id
	}
	if id, ok := byName["Priya Nair"]; ok {
		fmt.Println("Priya Nair has ID", id)
	}

	fmt.Println("\n=== Group by department: map[string][]Employee ===")
	byDept := make(map[string][]Employee)
	totals := make(map[string]float64)
	for _, emp := range byID {
		byDept[emp.Department] = append(byDept[emp.Department], emp)
		totals[emp.Department] += emp.Salary
	}
	depts := make([]string, 0, len(byDept))
	for d := range byDept {
		depts = append(depts, d)
	}
	sort.Strings(depts)
	for _, d := range depts {
		fmt.Printf("%-12s %d employee(s), total ₹%.0f\n", d, len(byDept[d]), totals[d])
	}
}
