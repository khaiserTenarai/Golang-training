package main

import "fmt"

func main() {
	var inputVal int

	fmt.Print("Enter a number: ")
	fmt.Scan(&inputVal)

	originalVal := inputVal
	reversedVal := 0

	for inputVal > 0 {
		remainder := inputVal % 10
		reversedVal = reversedVal*10 + remainder
		inputVal = inputVal / 10
	}

	if originalVal == reversedVal {
		fmt.Printf("%d is a palindrome.\n", originalVal)
	} else {
		fmt.Printf("%d is not a palindrome.\n", originalVal)
	}
}
