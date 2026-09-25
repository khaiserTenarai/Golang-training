package main

import "fmt"

//Declaring variadic function

// variadic function means  It can take 0 to n args

func TotalSalary(salaries ...int64) (total int64) {
	for _, sal := range salaries {
		total += sal
	}
	return

}

func main() {
	// call the function
	fmt.Println(TotalSalary(100000, 200000, 300000))
	fmt.Println(TotalSalary(100000, 200000, 300000, 500000))

	fmt.Println(TotalSalary(100000, 200000, 300000, 7899557))
}
