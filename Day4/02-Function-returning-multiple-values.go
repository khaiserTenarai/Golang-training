package main

import "fmt"

func main(){

	//call function
	fmt.Println(calculator(10,20))
	fmt.Println(calculator(10,0))

}


// shortend == (a,b int)
func calculator(a , b int )(sum int , diff int , prod int , quotient int ){
	sum  = a+b
	diff = a-b
	prod = a*b

	if b != 0{
	quotient = a/b
	}else{
		
		fmt.Println("can not perform division by 0")
	}
	return

}