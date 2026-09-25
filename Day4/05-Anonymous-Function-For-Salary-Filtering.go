package main

import "fmt"

func main() {

	//create one anonymous function and store inside isHigh variable
	isHigh := func(sal int64) bool{
		if sal > 40000 {
			return true
		}
		return false
	}

	//let the compiler decide the array size based on my added elements
	salaries := [...]int64{
		10000,
		20000,
		30000,
		450000,
		34000,
		4500,
		45098}

	//filter the salaries

	for index, sal := range salaries {
		if isHigh(sal) {
			fmt.Println("Person ", index+1)
		}
	}
}
