package main

import "fmt"

func main() {
	name := "Sanskriti"

	indianName := "संस्‍कृति"

	fmt.Println("English name:", name)
	fmt.Println("Bytes:", len(name))
	fmt.Println("Runes:", len([]rune(name)))

	fmt.Println("\nIndian name:", indianName)
	fmt.Println("Bytes:", len([]byte(indianName)))
	fmt.Println("Runes:", len([]rune(indianName)))

	fmt.Println("\nCharacters using runes:")

	for _, r := range indianName {
		fmt.Printf("%c ", r)
	}

	fmt.Println()
}