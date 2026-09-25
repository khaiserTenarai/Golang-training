package main

import "fmt"

// Value receiver
func changeByValue(number int) {
	number = 100
}

// Pointer receiver
func changeByPointer(number *int) {
	*number = 100
}

func main() {
	number1 := 50
	number2 := 50

	fmt.Println("Before changeByValue:", number1)

	changeByValue(number1)

	fmt.Println("After changeByValue:", number1)

	fmt.Println()

	fmt.Println("Before changeByPointer:", number2)

	changeByPointer(&number2)

	fmt.Println("After changeByPointer:", number2)
}