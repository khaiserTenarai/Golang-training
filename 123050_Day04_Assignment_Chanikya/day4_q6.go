package main

import "fmt"

func IdGenerate() func() int {
	a := 0
	return func() int {
		a++
		return a
	}
}
func main() {
	GenerateId := IdGenerate()
	fmt.Println(GenerateId())
	fmt.Println(GenerateId())
	fmt.Println(GenerateId())
}
