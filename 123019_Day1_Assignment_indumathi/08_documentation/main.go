package main

import (
	"fmt"
	"godocs/shapes"
)

func main() {
	s := shapes.Rectangle{Width: 7, Height: 4}

	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}
