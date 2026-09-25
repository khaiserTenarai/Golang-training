package main

import (
	"errors"
	"fmt"
)

func calculateSquare(num int) (int, error){
	if num < 0{
		return 0, errors.New("negative numbers not allowed")
	}
	return num * num, nil
}

func main(){

	fmt.Println("\n********************************************")
	fmt.Println("3. Create functions returning (result, error).")
	fmt.Println("**********************************************")

	var val int
	fmt.Print("\nEnter an number to square: ")
	fmt.Scan(&val)

	sqResult, sqErr := calculateSquare(val)

	if sqErr != nil {
		fmt.Println("Error:", sqErr)
	} else {
		fmt.Println("Square Result:", sqResult)
	}
}
