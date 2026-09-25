package main

import "fmt"

func main() {
	a := 10
	fmt.Println("before modification value a=", a)
	p := &a
	*p = 20
	fmt.Println("after modification value a=", a)
}
