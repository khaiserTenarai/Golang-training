package main

import "fmt"

func main() {
	x := 99 

	fmt.Println("Outer x before block:", x)

	{ 
		x := 10 
		fmt.Println("Inner x inside block :", x)
	}

	fmt.Println("Outer x after block :", x)
}