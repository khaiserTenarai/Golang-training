package main

import (
	"errors"
	"fmt"
)

func div(a, b float64) (result float64, err error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

func main() {
	fmt.Println(div(2, 0))
}
