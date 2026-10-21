package main

import "fmt"

// Task 10: Array Pagination Logic Arithmetic
// Computes correct bounds safe slices for dataset parsing.
func PaginateSlice(items []string, page, pageSize int) []string {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	
	start := (page - 1) * pageSize
	if start > len(items) {
		return []string{}
	}
	
	end := start + pageSize
	if end > len(items) {
		end = len(items)
	}
	
	return items[start:end]
}

func main() {
	data := []string{"A", "B", "C", "D", "E"}
	paged := PaginateSlice(data, 2, 2)
	fmt.Printf("[Task 10] Paginated array block subset: %v\n", paged)
}
