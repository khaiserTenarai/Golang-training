package main

import "fmt"

func cal(a, b int) (int, float32) {
	sum := a + b
	diff := a - b
	return sum, float32(diff)
}
func main() {
	fmt.Println(cal(2, 3))
}
