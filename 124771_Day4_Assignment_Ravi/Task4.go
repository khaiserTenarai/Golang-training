package main

import "fmt"

func add(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {

	fmt.Println(add(1, 2, 3))
}
