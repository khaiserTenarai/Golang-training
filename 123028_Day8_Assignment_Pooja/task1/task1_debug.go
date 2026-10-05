package main

import "fmt"

func main() {
	numbers := []int{}
	if len(numbers) > 0 {
		fmt.Println("First element:", numbers[0])
	} else {
		fmt.Println("Slice is empty")
	}
}