package main

import (
	"fmt"
	"strings"
)

// Task 11: Dynamic Slice Filtering Query Lookup
// Implements functional element screening matching parameters.
func FilterSlice(items []string, query string) []string {
	var result []string
	for _, item := range items {
		if strings.Contains(strings.ToLower(item), strings.ToLower(query)) {
			result = append(result, item)
		}
	}
	return result
}

func main() {
	list := []string{"Software Engineer", "Product Manager", "Architect"}
	matches := FilterSlice(list, "engineer")
	fmt.Printf("[Task 11] Filtering lookup output elements: %v\n", matches)
}
