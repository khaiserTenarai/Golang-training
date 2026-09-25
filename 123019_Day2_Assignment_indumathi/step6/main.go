package main

import "fmt"

func main() {
	age := 25
	salary := 40000

	fmt.Println("Age:", age)
	fmt.Println("Salary:", salary)

	fmt.Println("Age >= 18:", age >= 18)
	fmt.Println("Salary > 30000:", salary > 30000)

	fmt.Println("Eligible:", age >= 18 && salary > 30000)
	fmt.Println("Not eligible:", age < 18 || salary <= 30000)
	fmt.Println("Age is not 30:", age != 30)
}
