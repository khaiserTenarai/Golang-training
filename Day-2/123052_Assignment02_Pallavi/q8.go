package main

import "fmt"

func main() {
	name := "पल्लवी"

	fmt.Println("Name:", name)
	fmt.Println("Number of bytes:", len(name))

	runes := []rune(name)

	fmt.Println("Number of runes:", len(runes))

	for _, r := range runes {
		fmt.Printf("%c ", r)
	}
}