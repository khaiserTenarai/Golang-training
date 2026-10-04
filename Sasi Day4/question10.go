package main

import "fmt"

func modifyValue(num int) {
	num = 100
}

func modifyPointer(num *int) {
	*num = 100
}

func main() {
	value := 20

	modifyValue(value)
	fmt.Println(value)

	modifyPointer(&value)
	fmt.Println(value)
}