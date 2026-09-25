package main

import "fmt"

func main(){

	fmt.Println("\n*********************************************")
	fmt.Println("2. Create a function returning multiple values.")
	fmt.Println("***********************************************")

	addition, subtraction, multiplication, division := calculator(2, 3)
	fmt.Println("Addition:", addition, "Subtraction:", subtraction, "Multiplication:", multiplication, "Division:", division)

}

func calculator(a, b int) (int, int, int, int){
	add := a + b
	sub := a - b
	mul := a * b
	div := a / b

	return add, sub, mul, div
}
