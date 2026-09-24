package main

import (
	"fmt"
	"strconv"
)


func main(){

	fmt.Println("\n****************************************************")
	fmt.Println("Type-conversion program for string → integer → float. ")
	fmt.Println("******************************************************")

	// String to Integer conversion
	strVal := "123"
	fmt.Println("Original String Value: ", strVal)
	intVal, err := strconv.Atoi(strVal)
	if err != nil{
		fmt.Println("Error while converting to Integer")
		return
	}
	fmt.Println("Converted Integer Value: ", intVal)

	// Integer to Float conversion
	floatVal := float64(intVal)
	fmt.Println("Converted Float64 Value:", floatVal)

	fmt.Println("Thank You")

}