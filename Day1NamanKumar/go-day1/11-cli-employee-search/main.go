// Command employeesearch looks up an employee by ID from in-memory data.
//
//	go run . 102
package main

import (
	"fmt"
	"os"
	"strconv"
)

// Employee holds employee details.
type Employee struct {
	ID         int
	Name       string
	Department string
	Email      string
	Salary     float64
}

// employees is the in-memory "database", keyed by ID for O(1) lookup.
var employees = map[int]Employee{
	101: {101, "Asha Rao", "Engineering", "asha@corp.in", 1200000},
	102: {102, "Vikram Iyer", "Finance", "vikram@corp.in", 950000},
	103: {103, "Meera Nair", "HR", "meera@corp.in", 800000},
	104: {104, "Rahul Sharma", "Engineering", "rahul@corp.in", 1500000},
}

// FindByID returns the employee and whether it was found.
func FindByID(id int) (Employee, bool) {
	e, ok := employees[id]
	return e, ok
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: employeesearch <employee-id>")
		os.Exit(1)
	}
	id, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %q is not a valid employee ID\n", os.Args[1])
		os.Exit(1)
	}
	e, ok := FindByID(id)
	if !ok {
		fmt.Printf("No employee found with ID %d\n", id)
		os.Exit(1)
	}
	fmt.Println("Employee Details")
	fmt.Println("----------------")
	fmt.Printf("ID         : %d\n", e.ID)
	fmt.Printf("Name       : %s\n", e.Name)
	fmt.Printf("Department : %s\n", e.Department)
	fmt.Printf("Email      : %s\n", e.Email)
	fmt.Printf("Salary     : %.2f\n", e.Salary)
}
