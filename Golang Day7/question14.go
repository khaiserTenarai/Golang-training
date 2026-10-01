package main

import (
	"fmt"
	"strconv"
)

var Total int

func ProcessData(input string, values []int) string {
	count := 0
	result := ""

	num, _ := strconv.Atoi(input)
	
	Total = Total + num

	var myMap map[string]int
	myMap["key"] = num

	for i := 1; i <= len(values); i++ {
		result = result + strconv.Itoa(values[i])
	}

	if num > 10 {
		return result
	}

	average := Total / count
	fmt.Printf("Average: %d\n", average)

	return result
}

func main() {
	ProcessData("bad_input", []int{5, 10, 15})
}


/*Ignored Error Handling: Using the blank identifier _ to catch the strconv.Atoi error is dangerous. Because "bad_input" is passed in main, the conversion fails, the error is silenced, and num defaults to 0 without warning.

Nil Map Panic: var myMap map[string]int declares a map but does not allocate memory for it. Assigning myMap["key"] = num will immediately crash the program with a runtime panic. It should be initialized using make(map[string]int).

Division by Zero Panic: The count variable is set to 0 and is never changed. When the code reaches Total / count, it will crash with a divide-by-zero panic.

Index Out of Bounds: The loop condition uses i <= len(values). Because Go slices are zero-indexed, a slice of length 3 only has indices 0, 1, and 2. Trying to access values[3] will cause a panic.

Skipped First Element: The loop counter starts at i := 1. This means it completely misses the first element of the slice (values[0]).

Inefficient String Concatenation: Using result = result + ... inside a loop forces Go to allocate new memory and copy the string on every single iteration. For building strings in loops, strings.Builder is much more efficient.

Global State Mutation: Modifying the global variable Total inside a function is a bad practice. It makes the function unpredictable and will cause race conditions if the function is ever called by multiple goroutines at the same time.

Broken Logic: The count variable is clearly intended to track the number of items for the average calculation, but it is never actually incremented inside the loop.

Dead Code/Unused Data: myMap is populated with a key-value pair but is never read, returned, or used anywhere else in the application.

Violates Single Responsibility Principle: The ProcessData function tries to do too much at once. It parses strings, updates global variables, builds strings, handles business logic (thresholds), and prints to the console. These should be split into smaller, focused functions.*/