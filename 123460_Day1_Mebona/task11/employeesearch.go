package main

import (
	"fmt"
	"os"
)

func main() {
	
	employees := [][]string{
		{"101", "Alice (Developer)"},
		{"102", "Bob (Manager)"},
		
	}

	
	if len(os.Args) < 2 {
		fmt.Println("Error")
		return 
	}

	searchID := os.Args[1]
	found := false

	
	for _, emp := range employees {
		
		if emp[0] == searchID {
			
			fmt.Printf("Employee Found: %s\n", emp[1])
			found = true
			break 
		}
	}

	if !found {
		fmt.Printf("Employee with ID %s not found.\n", searchID)
	}

	/* 
		PS C:\Users\mebona.joseph\Desktop\day1_go> go run employeesearch.go 102
		Employee Found: Bob (Manager)
	*/
}