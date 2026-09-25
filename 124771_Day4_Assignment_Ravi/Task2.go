package main

import "fmt"

func calculate(a, b, c int) (int, int, int) {
	sum := a + b
	difference := a - b
	multiply := a * c
	return sum, difference, multiply
}

func main() {

	fmt.Println(calculate(1, 2, 3))

}
