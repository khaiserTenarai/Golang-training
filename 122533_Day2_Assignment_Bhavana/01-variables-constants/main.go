// 1. Create a program demonstrating Go variables and constants.

package main

import "fmt"

const Pi = 3.14159

func main() {
	var name string = "Ravi"
	var age int = 28
	city := "Bangalore" 

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("City:", city)
	fmt.Println("Value of Pi:", Pi)
}
