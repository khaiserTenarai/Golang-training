package main

import "fmt"

func main() {
	name := "Pallavi"

	fmt.Println("Outside:", name)

	if true {
		name := "Swathi"

		fmt.Println("Inside:", name)
	}

	fmt.Println("Outside again:", name)
}