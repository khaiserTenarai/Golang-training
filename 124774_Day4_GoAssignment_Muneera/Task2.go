package main

import "fmt"

func calculate(a, b int) (int, int, int, float64, int) {
	add := a + b
	sub := a - b
	mul := a * b
	div := float64(a) / float64(b)
	mod := a % b

	return add, sub, mul, div, mod
}

func main() {
	a := 9
	b := 2

	add, sub, mul, div, mod := calculate(a, b)
	fmt.Println("Addition", add)
	fmt.Println("Subtraction", sub)
	fmt.Println("Multiplication", mul)
	fmt.Println("Division", div)
	fmt.Println("Modulo Division(Remainder)", mod)

}
