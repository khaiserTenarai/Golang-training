package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run assignment_task11.go <employee_id>")
		return
	}

	ids := [3]string{"101", "102", "103"}
	names := [3]string{"Vimal", "Aditya", "Mason"}
	roles := [3]string{"Go Developer", "Product Manager", "DevOps Engineer"}

	searchID := os.Args[1]
	found := false

	for index, id := range ids {
		if id == searchID {
			fmt.Println("--- Employee Found ---")
			fmt.Println("ID:", ids[index])
			fmt.Println("Name:", names[index])
			fmt.Println("Role:", roles[index])
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Employee ID not found:", searchID)
	}
}