package main

import "fmt"

func changeSalary(salary *float64) {
	*salary = 40000
}

func main() {
	salary := 30000.0
	fmt.Println("Before Modification", salary)
	changeSalary(&salary)
	fmt.Println("After Modification:", salary)

}
