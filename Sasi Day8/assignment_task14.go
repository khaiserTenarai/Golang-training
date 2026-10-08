package main

import (
	"errors"
	"fmt"
	"strings"
	"sync/atomic"
)

type Employee1 struct {
	ID     int
	Name   string
	Salary float64
}

// // Issue 1 & 2: Incorrect signature and missing error handling
// func DivideBonus(total float64, count int) float64 {
// 	return total / float64(count)
// }

// // Issue 3: Inefficient string concatenation in loop
// func BuildEmployeeList(names []string) string {
// 	result := ""
// 	for _, name := range names {
// 		result += name + ", "
// 	}
// 	return result
// }

// // Issue 4: Exported function missing comment & poor variable names
// func Calc(a float64, b float64) float64 {
// 	return a * b
// }

// // Issue 5: Modifying copy instead of pointer
// func UpdateName(e Employee, newName string) {
// 	e.Name = newName
// }

// // Issue 6: Data race potential without synchronization
// var globalCounter = 0

// func IncrementCounter() {
// 	go func() {
// 		globalCounter++
// 	}()
// }

// // Issue 7: Ignoring errors explicitly
// func PrintDetails(e Employee) {
// 	formatted, _ := FormatName(e.Name)
// 	fmt.Println(formatted)
// }

// func FormatName(name string) (string, error) {
// 	if name == "" {
// 		return "", fmt.Errorf("empty name")
// 	}
// 	return strings.TrimSpace(name), nil
// }

// // Issue 8: Slice bounds / index panic vulnerability
// func GetFirstElement(items []string) string {
// 	return items[0]
// }

// // Issue 9: Bad format directive in Printf
// func PrintEmployeeInfo(e Employee) {
// 	fmt.Printf("Employee ID: %d, Name: %s\n", e.ID)
// }

// // Issue 10: Unreachable code
// func GetStatus(active bool) string {
// 	if active {
// 		return "Active"
// 	} else {
// 		return "Inactive"
// 	}
// 	return "Unknown"
// }

// DivideBonus safely divides total bonus amount among employee count.
func DivideBonus(total float64, count int) (float64, error) {
	if count <= 0 {
		return 0, errors.New("count must be greater than zero")
	}
	return total / float64(count), nil
}

// BuildEmployeeList concatenates employee names into a comma-separated string efficiently.
func BuildEmployeeList(names []string) string {
	return strings.Join(names, ", ")
}

// CalculateTotalSalary calculates the total salary based on base pay and multiplier.
func CalculateTotalSalary(baseSalary float64, multiplier float64) float64 {
	return baseSalary * multiplier
}

// UpdateName mutates the employee's name using a pointer reference.
func UpdateName(e *Employee1, newName string) {
	if e != nil {
		e.Name = newName
	}
}

var globalCounter int64

// IncrementCounter safely increments global counter using atomic operations.
func IncrementCounter() {
	atomic.AddInt64(&globalCounter, 1)
}

// PrintDetails formats and prints employee details while handling errors.
func PrintDetails(e Employee1) error {
	formatted, err := FormatName(e.Name)
	if err != nil {
		return fmt.Errorf("failed to format name: %w", err)
	}
	fmt.Println(formatted)
	return nil
}

func FormatName(name string) (string, error) {
	if strings.TrimSpace(name) == "" {
		return "", errors.New("empty name")
	}
	return strings.TrimSpace(name), nil
}

// GetFirstElement safely returns the first item or an error if empty.
func GetFirstElement(items []string) (string, error) {
	if len(items) == 0 {
		return "", errors.New("slice is empty")
	}
	return items[0], nil
}

// PrintEmployeeInfo correctly formats employee output parameters.
func PrintEmployeeInfo(e Employee1) {
	fmt.Printf("Employee ID: %d, Name: %s\n", e.ID, e.Name)
}

// GetStatus returns employee active status string cleanly.
func GetStatus(active bool) string {
	if active {
		return "Active"
	}
	return "Inactive"
}

func main() {
	emp := Employee1{ID: 101, Name: "Vimal", Salary: 500000}

	// 1. Demonstrate pointer modification
	UpdateName(&emp, "Vimal Aditya")

	// 2. Print formatted employee info
	PrintEmployeeInfo(emp)

	// 3. Demonstrate safe bonus calculation
	bonus, err := DivideBonus(50000, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Bonus per employee: %.2f\n", bonus)
	}

	// 4. Demonstrate status check
	fmt.Println("Status:", GetStatus(true))
}