package main

import "fmt"

func main() {
	age := 23
	salary := 50000
	if age >= 18 {
		fmt.Println("Age is", age)
	}
	fmt.Println(age >= 18 && salary > 30000)
	fmt.Println(age < 18 || salary > 30000)
	fmt.Println(!(age < 18))
}
