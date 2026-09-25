package main

import "fmt"

func changeSalary(salary *int) {
	*salary = 60000
}

func main() {
	salary := 50000

	fmt.Println("Before:", salary)

	changeSalary(&salary)

	fmt.Println("After:", salary)
}