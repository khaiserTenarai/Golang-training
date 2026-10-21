package main

import (
	"fmt"
	"sort"
)

// Task 12: In-Memory Sorting Implementation
// Orders data collections deterministically based on property fields.
func main() {
	names := []string{"Charlie", "Alice", "Bob"}
	sort.Strings(names)
	fmt.Printf("[Task 12] Sorted in-memory collections array: %v\n", names)
}
