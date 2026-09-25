package main

import (
	"fmt"
	"godocs/shapes" 
)

func main() {
	r := shapes.Rectangle{Width: 5, Height: 3}

	fmt.Println("Area:", r.Area())
	fmt.Println("Perimeter:", r.Perimeter())
}