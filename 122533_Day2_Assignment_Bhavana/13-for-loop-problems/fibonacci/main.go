// 13b. Fibonacci
//
// Solve using a for loop: print the first N numbers of the Fibonacci series.

package main

import "fmt"

func main() {
	n := 10
	first, second := 0, 1

	fmt.Print("Fibonacci series: ")
	for i := 0; i < n; i++ {
		fmt.Print(first, " ")
		first, second = second, first+second
	}
	fmt.Println()
}
