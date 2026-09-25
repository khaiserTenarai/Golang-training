package main

import "fmt"

func modifyValue(a int) {
	a = 20
}

func modifyPointer(a *int) {
	*a = 20
}

func main() {
	a := 10

	modifyValue(a)
	fmt.Println("After value:", a)

	modifyPointer(&a)
	fmt.Println("After pointer:", a)
}

/*Value: A copy of a is passed, so the original value is not modified.

Pointer: The address of a is passed, so the original value can be modified.*/
