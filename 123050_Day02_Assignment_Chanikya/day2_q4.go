package main

import (
	"fmt"
	"strconv"
)

func main() {
	st := "12345"
	fmt.Printf("Type: %T\n", st)
	n, err := strconv.Atoi(st)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return
	}
	fmt.Printf("Type: %T\n", n)
	f := float64(n)
	fmt.Printf("Type: %T\n", f)
}
