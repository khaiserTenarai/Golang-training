package main

import "fmt"

func main() {
	name := "ಇಂದುಮತಿ"

	fmt.Println("Name:", name)
	fmt.Println("Number of bytes:", len([]byte(name)))
	fmt.Println("Number of runes:", len([]rune(name)))

	fmt.Println("Bytes:", []byte(name))
	fmt.Println("Runes:", []rune(name))
}