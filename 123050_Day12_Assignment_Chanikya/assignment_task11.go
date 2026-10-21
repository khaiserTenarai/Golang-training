package main

import (
	"fmt"
	"strings"
)

// Task 11: Dynamic slice filtering query lookup
// Filtering matching entities functionally from data collections.

type Worker struct {
	Name string
	Dept string
}

func FilterWorkers(workers []Worker, department string) []Worker {
	var result []Worker
	for _, w := range workers {
		if strings.EqualFold(w.Dept, department) {
			result = append(result, w)
		}
	}
	return result
}

func main() {
	list := []Worker{{Name: "Alice", Dept: "IT"}, {Name: "Bob", Dept: "HR"}}
	fmt.Println("Task 11 - Filter IT:", FilterWorkers(list, "IT"))
}
