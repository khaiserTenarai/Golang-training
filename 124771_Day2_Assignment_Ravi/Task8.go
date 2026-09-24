package main

import "fmt"

func main() {
	name := "Ravi_Ranjan"

	fmt.Println("Name:", name)
	fmt.Println("Bytes:", len(name))
	fmt.Println("Runes:", len([]rune(name)))
}
