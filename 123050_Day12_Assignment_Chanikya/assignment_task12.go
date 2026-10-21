package main

import (
	"fmt"
	"sort"
)

// Task 12: In-Memory Sorting Implementation
// Dynamic sort structure mapping standard library sort algorithms.

type Product struct {
	Name  string
	Price float64
}

func SortProductsByPrice(products []Product) {
	sort.Slice(products, func(i, j int) bool {
		return products[i].Price < products[j].Price
	})
}

func main() {
	p := []Product{{"Tablet", 400}, {"Phone", 800}}
	SortProductsByPrice(p)
	fmt.Println("Task 12 - Sorted products array:", p)
}
