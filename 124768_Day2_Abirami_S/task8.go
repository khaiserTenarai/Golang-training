package main

import "fmt"

func main() {
	name := "Abirami"
	runes := []rune(name)
	fmt.Println("Name: ", name)
	fmt.Println("Bytes: ", len(name))
	fmt.Println("Runes: ", len(runes))
}
