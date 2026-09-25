package main

import "fmt"

func modifyValue(num *int) {
	*num = 200
}

func main() {
	value := 50
	
	fmt.Println(value)
	
	modifyValue(&value)
	
	fmt.Println(value)
}