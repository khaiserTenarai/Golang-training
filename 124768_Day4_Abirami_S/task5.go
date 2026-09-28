package main

import "fmt"

func main() {
	var salary float64
	filterSalary := func(sal int) bool {
		return sal > 50000
	}
	fmt.Println("Enter salary: ")
	fmt.Scan(&salary)
	fmt.Println(filterSalary(int(salary)))
}
