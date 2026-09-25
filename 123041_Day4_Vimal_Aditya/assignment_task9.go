package main

import "fmt"

func main(){

	fmt.Println("\n****************************************")
	fmt.Println("9. Demonstrate pointer-based modification.")
	fmt.Println("******************************************")

	x := 10

	p := &x  // this is the address of the x

	fmt.Println(p) // this will print the address

	*p = 20 // * represnts the value, that is changing the value prsent in the address from 10 to 20

	fmt.Println(x)

}