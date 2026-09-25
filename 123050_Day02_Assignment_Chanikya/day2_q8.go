package main

import "fmt"

func main() {
	name := "चाणक्य"

	fmt.Println("Name:", name)
	byte_len := len(name)
	fmt.Println("Number of bytes:", byte_len)

	runeSlice := []rune(name)
	runes_len := len(runeSlice)
	fmt.Println("Number of runes:", runes_len)

	for _, l := range runeSlice {
		fmt.Printf("%c ", l)
	}
}
