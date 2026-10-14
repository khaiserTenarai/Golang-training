package main

import (
	"fmt"
)

func Paginate(items []string, limit int, offset int) []string {
	if offset > len(items) {
		return []string{}
	}
	end := offset + limit
	if end > len(items) {
		end = len(items)
	}
	return items[offset:end]
}

func main() {
	data := []string{"Emp1", "Emp2", "Emp3", "Emp4", "Emp5"}
	paged := Paginate(data, 2, 1) // Limit 2, Offset 1
	fmt.Printf("Paginated Result: %v
", paged)
}
