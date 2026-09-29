// 9. Demonstrate pointer-based modification.

package main

import "fmt"

func doubleSalary(salary *int) {
	*salary = *salary * 2
}

func main() {
	salary := 20000
	fmt.Println("Before:", salary)

	doubleSalary(&salary)
	fmt.Println("After:", salary)
}
