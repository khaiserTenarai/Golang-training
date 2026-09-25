package main

import "fmt"

func changeValue(x int) {
	x = 100
}

func changePointer(x *int) {
	*x = 100
}

func main() {
	number1 := 50
	number2 := 50

	changeValue(number1)

	fmt.Println("After value:", number1)

	changePointer(&number2)

	fmt.Println("After pointer:", number2)
}