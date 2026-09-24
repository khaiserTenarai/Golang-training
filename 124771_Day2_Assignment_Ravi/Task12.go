package main

import "fmt"

func main() {
	var choice int
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	a, b := 10, 5

	switch choice {
	case 1:
		fmt.Println("Result:", a+b)
	case 2:
		fmt.Println("Result:", a-b)
	case 3:
		fmt.Println("Result:", a*b)
	case 4:
		fmt.Println("Result:", a/b)
	default:
		fmt.Println("Invalid choice")
	}
}
