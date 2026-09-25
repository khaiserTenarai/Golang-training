package main

import (
	"fmt"
	"strconv"
)

// 4. Type Conversion Program (String $\rightarrow$ Integer $\rightarrow$ Float)

func main() {

	n1 := "100"

	// string -> int
	n2, err := strconv.Atoi(n1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}


	n3 := float64(n2)

	fmt.Println("String       :", n1)
	fmt.Println("int  :", n2)
	fmt.Println("Float64      :", n3)

}
