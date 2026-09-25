package main

import "fmt"

func main() {
	x := 10
	p := &x
	fmt.Println(*p)
	*p = 20
	fmt.Println(*p)
}
