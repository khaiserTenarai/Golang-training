package main

import "fmt"

func main() {

	fmt.Println("\n********************************************************")
	fmt.Println("8. Demonstrate bytes vs runes using Indian/Unicode names. ")
	fmt.Println("**********************************************************")

	var name string

	// विमल आदित्य

	fmt.Print("Enter a name (in Hindi/): ")
	fmt.Scanln(&name)

	runes := []rune(name)

	byteCount := len(name)
	runeCount := len(runes)

	fmt.Println("\n--- Bytes vs Runes Analysis ---")
	fmt.Println("Name            :", name)
	fmt.Println("Number of Bytes :", byteCount)
	fmt.Println("Number of Runes :", runeCount)
	
	fmt.Print("Individual Runes: ")
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
	fmt.Println()
}