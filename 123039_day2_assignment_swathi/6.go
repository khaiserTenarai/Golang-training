
package main

import "fmt"

func main() {

	age := 25
	salary := 50000

	fmt.Println("===== COMPARISON OPERATORS =====")

	// Greater than
	fmt.Println("Age > 18:", age > 18)

	// Less than
	fmt.Println("Age < 30:", age < 30)

	// Greater than or equal to
	fmt.Println("Age >= 25:", age >= 25)

	// Less than or equal to
	fmt.Println("Age <= 25:", age <= 25)

	// Equal to
	fmt.Println("Age == 25:", age == 25)

	// Not equal to
	fmt.Println("Age != 30:", age != 30)

	fmt.Println("\n===== LOGICAL OPERATORS =====")

	// AND operator
	fmt.Println("Age > 18 AND Salary > 40000:",
		age > 18 && salary > 40000)

	// OR operator
	fmt.Println("Age < 18 OR Salary > 40000:",
		age < 18 || salary > 40000)

	// NOT operator
	fmt.Println("NOT Age > 18:",
		!(age > 18))
}


