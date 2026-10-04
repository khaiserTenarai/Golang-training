package main

import "fmt"

func main() {
	name := "Sasikumar"

	fmt.Println("Outside:", name)

	if true {
		name := "Ajith"

		fmt.Println("Inside:", name)
	}

	fmt.Println("Outside again:", name)
}