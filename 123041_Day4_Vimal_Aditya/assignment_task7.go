package main

import "fmt"

func factorial(n int) int{
	if n <= 1 {
		return 1
	}
	return  n * factorial(n - 1)
}

func main(){

	fmt.Println("\n*************************************")
	fmt.Println("7. Implement factorial using recursion.")
	fmt.Println("***************************************")

	fact := factorial(5)
	fmt.Println(fact)

}