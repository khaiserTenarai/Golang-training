package view

import "fmt"

func ShowMenu() {

	fmt.Println("\n========== E-COMMERCE MANAGEMENT ==========")
	fmt.Println("1. Add Product")
	fmt.Println("2. Display Products")
	fmt.Println("3. Update Product")
	fmt.Println("4. Delete Product")
	fmt.Println("5. Increase Stock")
	fmt.Println("6. Decrease Stock")
	fmt.Println("7. Show Low Stock Products")
	fmt.Println("8. Exit")
	fmt.Println("===========================================")
}

func GetChoice() int {

	var choice int

	fmt.Print("Enter choice: ")
	fmt.Scanln(&choice)

	return choice
}
