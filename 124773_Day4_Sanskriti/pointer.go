
package main

import "fmt"

func increaseSalary(salary *float64) {
	*salary = *salary + 5000
}

func main() {

	salary := 30000.0

	fmt.Println("Before:", salary)

	increaseSalary(&salary)

	fmt.Println("After:", salary)
}