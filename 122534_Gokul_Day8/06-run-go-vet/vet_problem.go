package main

import "fmt"

func main() {
	name := "Anita"
	age := 28

	// BUG: %d expects a number, but "name" is a string. go vet's printf
	// checker catches exactly this kind of mismatch at build time.
	fmt.Printf("Employee %d is %d years old\n", name, age)
}
