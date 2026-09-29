package main

import "fmt"

func main() {

	var num int
	reverse := 0

	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	for num > 0 {

		digit := num % 10

		reverse = reverse*10 + digit

		num = num / 10
	}

	fmt.Println("Reverse:", reverse)
}