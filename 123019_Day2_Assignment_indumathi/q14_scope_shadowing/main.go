package main

import "fmt"

func main() {
	x := 10
	fmt.Printf("1. Outer block, x = %d\n", x)

	{
		x := 20 // Inner x shadows the outer x within this block
		fmt.Printf("2. Inner block, x = %d\n", x)
	}

	fmt.Printf("3. Outer block again, x = %d\n", x)
}