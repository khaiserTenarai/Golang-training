package main

import "fmt"

func main() {

	//local var that stores  value 10
	// a points towards  value 10
	a := 10
	fmt.Println("a value (Before modification): ", a)

	//store the address , where that 10 is present
	p := &a

	//print that address
	fmt.Println("address of memory where 10 stored :", p)
	fmt.Println("address of memory where 10 stored :", &a)

	//get that value 
	fmt.Println("value  :", *p)

	//get that value  (here * nullifies &)
	fmt.Println("value  :", *&a)

//after change
	*&a = 90
	fmt.Println(a)



}
