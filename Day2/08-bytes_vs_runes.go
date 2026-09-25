package main

import "fmt"

func main() {
	name := "गणेश"

	fmt.Println("Name:", name)

	// Number of bytes
	fmt.Println("Bytes:", len(name))

	// Number of runes
	runes := 0

	for range name {
		runes++
	}

	fmt.Println("Runes:", runes)
}
