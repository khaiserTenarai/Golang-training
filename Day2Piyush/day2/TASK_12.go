package main

import "fmt"

func main() {

	var choice int
	var a, b float64

	fmt.Println("----- Calculator Menu -----")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Println("5. Exit")

	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	switch choice {

	case 1:
		fmt.Print("Enter first number: ")
		fmt.Scan(&a)

		fmt.Print("Enter second number: ")
		fmt.Scan(&b)

		fmt.Println("Result:", a+b)

	case 2:
		fmt.Print("Enter first number: ")
		fmt.Scan(&a)

		fmt.Print("Enter second number: ")
		fmt.Scan(&b)

		fmt.Println("Result:", a-b)

	case 3:
		fmt.Print("Enter first number: ")
		fmt.Scan(&a)

		fmt.Print("Enter second number: ")
		fmt.Scan(&b)

		fmt.Println("Result:", a*b)

	case 4:
		fmt.Print("Enter first number: ")
		fmt.Scan(&a)

		fmt.Print("Enter second number: ")
		fmt.Scan(&b)

		if b == 0 {
			fmt.Println("Cannot divide by zero")
		} else {
			fmt.Println("Result:", a/b)
		}

	case 5:
		fmt.Println("Program exited.")

	default:
		fmt.Println("Invalid choice")
	}
}