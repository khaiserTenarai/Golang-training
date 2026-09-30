// Task 14: Variable scope and shadowing.
//
// Scope = where a variable can be used. Go uses BLOCK scope: a
// variable lives from its declaration to the closing } of its block.
// Shadowing = declaring a new variable with the same name in an inner
// block, which hides the outer one.
package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Package scope: visible in every function of this package.
var appName = "EmployeeApp"

// Exported (capitalised) package-level names are visible to OTHER packages too.
var Version = "v2.0.0"

func main() {
	fmt.Println("=== 1. Package scope ===")
	fmt.Println("appName from main:", appName)
	printAppName()

	fmt.Println("\n=== 2. Function scope ===")
	salary := 50000 // lives until the end of main
	fmt.Println("salary in main:", salary)
	// fmt.Println(bonus) // compile error: bonus only exists inside giveBonus
	giveBonus()

	fmt.Println("\n=== 3. Block scope ===")
	if salary > 40000 {
		level := "Senior" // exists only inside this if block
		fmt.Println("inside if, level =", level)
	}
	// fmt.Println(level) // compile error: undefined: level

	for i := 0; i < 2; i++ { // i exists only inside the loop
		fmt.Println("loop i =", i)
	}
	// fmt.Println(i) // compile error: undefined: i

	fmt.Println("\n=== 4. if / switch init statements ===")
	if n, err := strconv.Atoi("42"); err == nil {
		fmt.Println("n and err live only inside this if/else:", n)
	}

	fmt.Println("\n=== 5. Shadowing ===")
	x := 10
	fmt.Println("outer x =", x)
	{
		x := 20 // NEW variable that hides the outer x
		fmt.Println("  inner x =", x)
		x++
		fmt.Println("  inner x after ++ =", x)
	}
	fmt.Println("outer x is unchanged =", x)

	fmt.Println("\nShadowing the package variable:")
	appName := "LocalName" // hides package-level appName in main
	fmt.Println("in main  :", appName)
	printAppName() // still sees the package-level one

	fmt.Println("\n=== 6. The classic shadowing BUG ===")
	count := 0
	if true {
		count := 5 // meant count = 5, but := made a NEW variable
		_ = count
	}
	fmt.Println("buggy: count =", count, "(expected 5)")

	count = 0
	if true {
		count = 5 // = assigns to the existing variable
	}
	fmt.Println("fixed: count =", count)

	fmt.Println("\n=== 7. The error shadowing BUG ===")
	fmt.Println("buggy result:", loadBuggy())
	fmt.Println("fixed result:", loadFixed())

	fmt.Println("\n=== 8. Closures capture variables from the outer scope ===")
	counter := makeCounter()
	fmt.Println(counter(), counter(), counter())

	fmt.Println("\nTip: find shadowing with the shadow analyzer:")
	fmt.Println("  go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest")
	fmt.Println("  go vet -vettool=$(which shadow) ./...")
}

func printAppName() { fmt.Println("in printAppName:", appName) }

func giveBonus() {
	bonus := 5000 // function scope: invisible outside giveBonus
	fmt.Println("bonus in giveBonus:", bonus)
}

func validate() error { return errors.New("validation failed") }

// loadBuggy returns nil even though validate() failed.
func loadBuggy() error {
	var err error
	if true {
		_, err := strconv.Atoi("1") // := creates a new err in this block
		if err == nil {
			err = validate() // sets the INNER err only
		}
		_ = err
	}
	return err // the outer err is still nil
}

// loadFixed declares the value first, then uses = so the outer err is set.
func loadFixed() error {
	var err error
	if true {
		var n int
		n, err = strconv.Atoi("1")
		_ = n
		if err == nil {
			err = validate()
		}
	}
	return err
}

func makeCounter() func() int {
	count := 0 // survives between calls because the closure captures it
	return func() int {
		count++
		return count
	}
}
