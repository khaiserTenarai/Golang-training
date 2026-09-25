package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "42"

	i, _ := strconv.Atoi(s)
	fmt.Println("String to Int:", i)

	f := float64(i)
	fmt.Println("Int to Float :", f)

	s2 := fmt.Sprintf("%.2f", f)
	fmt.Println("Float to String:", s2)
}