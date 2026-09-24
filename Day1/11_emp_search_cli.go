package main

import (
	"fmt"
	"os"
)

type worker struct {
	id   string
	name string
	role string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("provide an id.")
		return
	}

	search := os.Args[1]

	people := []worker{
		{"1", "Gokul", "admin"},
		{"2", "bhavan", "sales"},
		{"3", "sachin", "support"},
	}

	for _, p := range people {
		if p.id == search {
			fmt.Println("id:", p.id)
			fmt.Println("name:", p.name)
			fmt.Println("role:", p.role)
			return
		}
	}

	fmt.Println("no one found with id:", search)
}

// Command to run
// go run 11_emp_search_cli.go 2
