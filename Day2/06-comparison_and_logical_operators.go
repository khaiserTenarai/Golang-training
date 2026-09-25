package main

import "fmt"

func main() {

	age := 25
	salary := 50000

	// Comparison Operators or Relational Operators
	fmt.Println("Comparison Operators:")

	fmt.Println("age > 18  :", age > 18)
	fmt.Println("age < 18  :", age < 18)
	fmt.Println("age == 25 :", age == 25)
	fmt.Println("age != 30 :", age != 30)
	fmt.Println("age >= 25 :", age >= 25)
	fmt.Println("age <= 25 :", age <= 25)

	// Logical Operators
	fmt.Println("\nLogical Operators:")

	fmt.Println("true && true :", true && true)
	fmt.Println("true && false:", true && false)

	fmt.Println("true || false:", true || false)
	fmt.Println("false || false:", false || false)

	fmt.Println("!true :", !true)
	fmt.Println("!false:", !false)

	// Combination of Comparison + Logical Operators
	fmt.Println("\nCombination:")

	fmt.Println("Age >= 18 AND Salary >= 30000:", age >= 18 && salary >= 30000)

	fmt.Println("Age < 18 OR Salary > 40000:", age < 18 || salary > 40000)

	fmt.Println("NOT (Age == 30):", !(age == 30))
}
