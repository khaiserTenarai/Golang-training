package main

import "fmt"

func main() {
	var n int

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	f := 0
	s := 1

	for i := 0; i < n; i++ {
		fmt.Println(f)

		temp := f + s
		f = s
		s = temp
	}
}
