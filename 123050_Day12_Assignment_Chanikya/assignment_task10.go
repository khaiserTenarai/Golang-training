package main

import "fmt"

// Task 10: Array pagination logic arithmetic
// Avoids memory overflow errors by subsetting slices mathematically based on page requests.

func Paginate(items []string, page, pageSize int) []string {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	start := (page - 1) * pageSize
	if start >= len(items) {
		return []string{}
	}
	end := start + pageSize
	if end > len(items) {
		end = len(items)
	}
	return items[start:end]
}

func main() {
	list := []string{"A", "B", "C", "D", "E"}
	fmt.Println("Task 10 - Page 2 size 2 output:", Paginate(list, 2, 2))
}
