package main

import "fmt"

func main() {
	var (
		i int
		f float64
		s string
		b bool
		arr [2]int
		sl []int
		m map[string]int
		p *int
	)

	fmt.Println("int:", i)
	fmt.Println("float:", f)
	fmt.Printf("string: %q\n", s)
	fmt.Println("bool:", b)
	fmt.Println("array:", arr)
	fmt.Println("slice:", sl)
	fmt.Println("map:", m)
	fmt.Println("pointer:", p)
}