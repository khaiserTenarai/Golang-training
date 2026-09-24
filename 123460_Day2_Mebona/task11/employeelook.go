package main

import (
	"fmt"
)

type Member struct {
	Name string
	Job  string
}

func main() {
	
	roster := map[int]Member{
		101: {"Alice", "Developer"},
		102: {"Bob", "Manager"},
		103: {"Charlie", "Designer"},
	}

	var searchID int
	fmt.Print("Enter Member ID: ")
	fmt.Scan(&searchID)

	
	if m, exists := roster[searchID]; exists {
		fmt.Printf("Employee: %s  Role: %s\n", m.Name, m.Job)
	} else {
		fmt.Println("Employee not found.")
	}
}