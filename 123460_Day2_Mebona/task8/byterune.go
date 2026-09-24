package main

import "fmt"

func main() {
	
	name := "मेबोना"

	fmt.Println("Name:", name)
	
	fmt.Println("Number of bytes:", len(name))

	runes := []rune(name)

	fmt.Println("Number of runes:", len(runes))

	fmt.Print("Characters (Code points): ")
	for _, r := range runes {
		fmt.Printf("%c ", r)
	}
	fmt.Println()
}