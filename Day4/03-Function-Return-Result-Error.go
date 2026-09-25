package main

import (
	"fmt"
	"errors"
)

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Can not divide by zero")
	}

	return a / b, nil

}

func main() {

	//Test case -1 what if I pass 0 to b var
	result, err := Divide(10, 0)

	fmt.Println("------------Test case -1 ----------------")
	fmt.Println("result :",result)
	fmt.Println("error :",err)

	//Test case -2 what if I pass non-zero to b
	fmt.Println("---------------Test case -2 ----------------")
	result, err = Divide(10, 2)
	fmt.Println("result :",result)
	fmt.Println("error :",err)

}
