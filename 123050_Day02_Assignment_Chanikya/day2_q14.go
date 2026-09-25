package main

import "fmt"

func main() {
	count := 10
	fmt.Println("Main start:", count)

	if true {
		count := 55
		fmt.Println("Inside if:", count)

	}

	fmt.Println("Main end:", count)
}
