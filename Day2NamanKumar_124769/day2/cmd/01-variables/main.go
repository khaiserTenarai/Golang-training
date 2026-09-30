// Task 1: Go variables and constants.
package main

import "fmt"

// Package-level variable: visible to every function in this package.
var companyName = "Cloud Native Pvt Ltd"

// Constants never change and are fixed at compile time.
const (
	MaxLeavesPerYear         = 24   // untyped constant: adapts to the type it is used with
	TaxRate          float64 = 0.10 // typed constant: always float64
)

// Level uses iota to create auto-incrementing constants, like an enum.
type Level int

const (
	Intern  Level = iota // 0
	Junior               // 1
	Senior               // 2
	Lead                 // 3
	Manager              // 4
)

func (l Level) String() string {
	return [...]string{"Intern", "Junior", "Senior", "Lead", "Manager"}[l]
}

func main() {
	fmt.Println("=== Ways to declare variables ===")

	// 1. var with explicit type
	var name string = "Anita Rao"
	// 2. var with type inference
	var age = 29
	// 3. short declaration (only inside functions)
	salary := 85000.50
	// 4. declare several at once
	var city, state = "Bengaluru", "Karnataka"
	// 5. grouped declaration
	var (
		empID    int  = 101
		isActive bool = true
	)
	// 6. declare now, assign later (starts at its zero value)
	var department string
	department = "Engineering"

	fmt.Printf("%-11s %-22v type: %T\n", "name", name, name)
	fmt.Printf("%-11s %-22v type: %T\n", "age", age, age)
	fmt.Printf("%-11s %-22v type: %T\n", "salary", salary, salary)
	fmt.Printf("%-11s %-22v type: %T\n", "city/state", city+", "+state, city)
	fmt.Printf("%-11s %-22v type: %T\n", "empID", empID, empID)
	fmt.Printf("%-11s %-22v type: %T\n", "isActive", isActive, isActive)
	fmt.Printf("%-11s %-22v type: %T\n", "department", department, department)
	fmt.Println("company    ", companyName, "(package-level)")

	// Variables can change...
	age = age + 1
	fmt.Println("\nAfter birthday, age =", age)

	// Swap two variables in one line (multiple assignment)
	a, b := 10, 20
	a, b = b, a
	fmt.Println("Swapped: a =", a, "b =", b)

	// The blank identifier _ discards a value you don't need
	_, second := "ignored", "kept"
	fmt.Println("Blank identifier kept:", second)

	fmt.Println("\n=== Constants ===")
	fmt.Println("MaxLeavesPerYear:", MaxLeavesPerYear)
	fmt.Println("TaxRate         :", TaxRate)
	var halfLeaves float64 = MaxLeavesPerYear / 2.0 // untyped const works as float64
	fmt.Println("Half leaves     :", halfLeaves)
	// MaxLeavesPerYear = 30  // compile error: cannot assign to a constant

	fmt.Println("\n=== iota (enumerated constants) ===")
	for l := Intern; l <= Manager; l++ {
		fmt.Printf("%d = %v\n", l, l)
	}
}
