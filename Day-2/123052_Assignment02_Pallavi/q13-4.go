package main

import "fmt"

func main() {
	var number int

	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	reverse := 0

	for number > 0 {
		digit := number % 10
		reverse = reverse*10 + digit
		number = number / 10
	}

	fmt.Println("Reverse:", reverse)
}