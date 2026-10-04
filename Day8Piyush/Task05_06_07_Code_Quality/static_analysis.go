// Task 7: Perform Static Analysis
// Install: go install honnef.co/go/tools/cmd/staticcheck@latest
// Run: staticcheck static_analysis.go
// Or use: go vet ./...
//
// This file demonstrates code patterns caught by static analysis tools.

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// Issue 1: Inefficient string concatenation in loop (should use strings.Builder)
	result := ""
	for i := 0; i < 100; i++ {
		result += fmt.Sprintf("item_%d,", i)
	}
	fmt.Println("Length:", len(result))

	// Issue 2: Compiling regex inside a loop (should compile once outside)
	names := []string{"Alice123", "Bob", "Charlie456"}
	for _, name := range names {
		re := regexp.MustCompile(`[0-9]+`) // compiled every iteration
		fmt.Println(name, "->", re.ReplaceAllString(name, ""))
	}

	// Issue 3: Using strings.Replace with -1 when strings.ReplaceAll exists
	text := "hello world hello world"
	fmt.Println(strings.Replace(text, "hello", "hi", -1))

	// Issue 4: Empty branch
	x := 10
	if x > 5 {
		// TODO: handle this case
	}

	// Issue 5: Unnecessary type conversion
	var a int = 10
	b := int(a) // a is already int
	fmt.Println(b)
}
