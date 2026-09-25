// Day 1 - Task 13: Build Automation
//
// See the Makefile in this folder for the run / build / test / fmt / vet
// targets.
package main

import "fmt"

func Square(n int) int {
	return n * n
}

func main() {
	fmt.Println("Square of 6 is", Square(6))
}
